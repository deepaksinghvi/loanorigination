package common

import (
	"context"
	"errors"
	"fmt"
	prom "github.com/m3db/prometheus_client_golang/prometheus"
	"github.com/opentracing/opentracing-go"
	"github.com/uber-go/tally"
	"github.com/uber-go/tally/prometheus"
	apiv1 "github.com/uber/cadence-idl/go/proto/api/v1"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/compatibility"
	"go.uber.org/cadence/encoded"
	"go.uber.org/cadence/workflow"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/grpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
)

const (
	defaultConfigFile       = "config/development.yaml"
	_cadenceClientName      = "cadence-client"
	_cadenceFrontendService = "cadence-frontend"
)

var (
	safeCharacters = []rune{'_'}

	sanitizeOptions = tally.SanitizeOptions{
		NameCharacters: tally.ValidCharacters{
			Ranges:     tally.AlphanumericRange,
			Characters: safeCharacters,
		},
		KeyCharacters: tally.ValidCharacters{
			Ranges:     tally.AlphanumericRange,
			Characters: safeCharacters,
		},
		ValueCharacters: tally.ValidCharacters{
			Ranges:     tally.AlphanumericRange,
			Characters: safeCharacters,
		},
		ReplacementCharacter: tally.DefaultReplacementCharacter,
	}
)

type CadenceHelper struct {
	Service            workflowserviceclient.Interface
	WorkerMetricScope  tally.Scope
	ServiceMetricScope tally.Scope
	Logger             *zap.Logger
	Config             Configuration
	Builder            *WorkflowClientBuilder
	DataConverter      encoded.DataConverter
	CtxPropagators     []workflow.ContextPropagator
	//workflowRegistries []registryOption
	//activityRegistries []registryOption
	Tracer opentracing.Tracer

	configFile string
}
type Configuration struct {
	DomainName      string                    `yaml:"domain"`
	ServiceName     string                    `yaml:"service"`
	HostNameAndPort string                    `yaml:"host"`
	TaskList        string                    `yaml:"tasklist"`
	Prometheus      *prometheus.Configuration `yaml:"prometheus"`
}

/*
type CadenceConfig struct {
	Builder       *WorkflowClientBuilder
	CadenceClient workflowserviceclient.Interface
	Logger        *zap.Logger
	Domain        string
	TaskList      string
	HostNamePart  string //host: "127.0.0.1:7833"
}*/

var LOCadenceHelper *CadenceHelper

func NewLOCadenceHelper(h *CadenceHelper) *CadenceHelper {
	LOCadenceHelper = h
	return h
}

// SetupServiceConfig setup the config for the sample code run
func (h *CadenceHelper) SetupServiceConfig() {
	if h.Service != nil {
		return
	}

	if h.configFile == "" {
		h.configFile = defaultConfigFile
	}
	// Initialize developer config for running Loan Origination
	configData, err := ioutil.ReadFile(h.configFile)
	if err != nil {
		panic(fmt.Sprintf("Failed to log config file: %v, Error: %v", defaultConfigFile, err))
	}

	if err := yaml.Unmarshal(configData, &h.Config); err != nil {
		panic(fmt.Sprintf("Error initializing configuration: %v", err))
	}

	// Initialize logger for running Loan Origination
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("Logger created.")
	h.Logger = logger
	h.ServiceMetricScope = tally.NoopScope
	h.WorkerMetricScope = tally.NoopScope

	if h.Config.Prometheus != nil {
		reporter, err := h.Config.Prometheus.NewReporter(
			prometheus.ConfigurationOptions{
				Registry: prom.NewRegistry(),
				OnError: func(err error) {
					logger.Warn("error in prometheus reporter", zap.Error(err))
				},
			},
		)
		if err != nil {
			panic(err)
		}

		h.WorkerMetricScope, _ = tally.NewRootScope(tally.ScopeOptions{
			Prefix:          "Worker_",
			Tags:            map[string]string{},
			CachedReporter:  reporter,
			Separator:       prometheus.DefaultSeparator,
			SanitizeOptions: &sanitizeOptions,
		}, 1*time.Second)

		// NOTE: this must be a different scope with different prefix, otherwise the metric will conflict
		h.ServiceMetricScope, _ = tally.NewRootScope(tally.ScopeOptions{
			Prefix:          "Service_",
			Tags:            map[string]string{},
			CachedReporter:  reporter,
			Separator:       prometheus.DefaultSeparator,
			SanitizeOptions: &sanitizeOptions,
		}, 1*time.Second)
	}
	h.Builder = NewBuilder(logger).
		SetHostPort(h.Config.HostNameAndPort).
		SetDomain(h.Config.DomainName).
		SetMetricsScope(h.ServiceMetricScope).
		SetDataConverter(h.DataConverter).
		SetTracer(h.Tracer).
		SetContextPropagators(h.CtxPropagators)
	service, err := h.Builder.BuildServiceClient()
	if err != nil {
		panic(err)
	}
	h.Service = service

	domainClient, _ := h.Builder.BuildCadenceDomainClient()
	_, err = domainClient.Describe(context.Background(), h.Config.DomainName)
	if err != nil {
		logger.Info("Domain doesn't exist", zap.String("Domain", h.Config.DomainName), zap.Error(err))
	} else {
		logger.Info("Domain successfully registered.", zap.String("Domain", h.Config.DomainName))
	}

	//h.workflowRegistries = make([]registryOption, 0, 1)
	//h.activityRegistries = make([]registryOption, 0, 1)
}

/*
func (c *CadenceConfig) Setup() {
	c.Logger = BuildLogger()
	c.Builder = NewBuilder(c.Logger).
		SetHostPort(h.Config.HostNameAndPort).
		SetDomain(h.Config.DomainName).
		SetMetricsScope(h.ServiceMetricScope).
		SetDataConverter(h.DataConverter).
		SetTracer(h.Tracer).
		SetContextPropagators(h.CtxPropagators)
	c.CadenceClient = BuildCadenceClient()
	c.Domain = Domain
	c.TaskList = TaskListName
}

*/

func BuildCadenceClient() workflowserviceclient.Interface {
	dispatcher := yarpc.NewDispatcher(yarpc.Config{
		Name: ClientName,
		Outbounds: yarpc.Outbounds{
			CadenceService: {Unary: grpc.NewTransport().NewSingleOutbound(HostPort)},
		},
	})
	if err := dispatcher.Start(); err != nil {
		panic("Failed to start dispatcher: " + err.Error())
	}

	clientConfig := dispatcher.ClientConfig(CadenceService)

	return compatibility.NewThrift2ProtoAdapter(
		apiv1.NewDomainAPIYARPCClient(clientConfig),
		apiv1.NewWorkflowAPIYARPCClient(clientConfig),
		apiv1.NewWorkerAPIYARPCClient(clientConfig),
		apiv1.NewVisibilityAPIYARPCClient(clientConfig),
	)
}

func BuildLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zapcore.InfoLevel)

	var err error
	logger, err := config.Build()
	if err != nil {
		panic("Failed to setup logger: " + err.Error())
	}

	return logger
}

/*
input json format
{
  "aadhaar_number": "123123123123",
  "applicant_name": "John Doe",
  "pan_number": "APWSDFDSF"
}
*/

func (c *CadenceHelper) StartLoanOriginationWorkflowExecution(workflowType, workflowID, runID string, input []byte) {
	executionTimeout := int32(60)
	closeTimeout := int32(60)

	req := shared.StartWorkflowExecutionRequest{
		Domain:     &c.Config.DomainName,
		WorkflowId: &workflowID, // loanWorkflowID
		WorkflowType: &shared.WorkflowType{
			Name: &workflowType,
		},
		TaskList: &shared.TaskList{
			Name: &c.Config.TaskList, // A task list is essentially a queue used to store an activity task until it is picked up by an available worker.
		},
		Input:                               input,
		ExecutionStartToCloseTimeoutSeconds: &executionTimeout,
		TaskStartToCloseTimeoutSeconds:      &closeTimeout,
		RequestId:                           &runID, //loanApplicationID
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resp, err := c.Service.StartWorkflowExecution(ctx, &req)

	if err != nil {
		c.Logger.Error("Failed to create workflow", zap.Error(err))
		panic("Failed to create workflow.")
	}

	c.Logger.Info("successfully started Loan Application workflow", zap.String("runID", resp.GetRunId()))
}

// WorkflowClientBuilder build client to cadence service
type WorkflowClientBuilder struct {
	hostPort       string
	dispatcher     *yarpc.Dispatcher
	domain         string
	clientIdentity string
	metricsScope   tally.Scope
	Logger         *zap.Logger
	ctxProps       []workflow.ContextPropagator
	dataConverter  encoded.DataConverter
	tracer         opentracing.Tracer
}

// NewBuilder creates a new WorkflowClientBuilder
func NewBuilder(logger *zap.Logger) *WorkflowClientBuilder {
	return &WorkflowClientBuilder{
		Logger: logger,
	}
}

// SetHostPort sets the hostport for the builder
func (b *WorkflowClientBuilder) SetHostPort(hostport string) *WorkflowClientBuilder {
	b.hostPort = hostport
	return b
}

// SetDomain sets the domain for the builder
func (b *WorkflowClientBuilder) SetDomain(domain string) *WorkflowClientBuilder {
	b.domain = domain
	return b
}

// SetClientIdentity sets the identity for the builder
func (b *WorkflowClientBuilder) SetClientIdentity(identity string) *WorkflowClientBuilder {
	b.clientIdentity = identity
	return b
}

// SetMetricsScope sets the metrics scope for the builder
func (b *WorkflowClientBuilder) SetMetricsScope(metricsScope tally.Scope) *WorkflowClientBuilder {
	b.metricsScope = metricsScope
	return b
}

// SetDispatcher sets the dispatcher for the builder
func (b *WorkflowClientBuilder) SetDispatcher(dispatcher *yarpc.Dispatcher) *WorkflowClientBuilder {
	b.dispatcher = dispatcher
	return b
}

// SetContextPropagators sets the context propagators for the builder
func (b *WorkflowClientBuilder) SetContextPropagators(ctxProps []workflow.ContextPropagator) *WorkflowClientBuilder {
	b.ctxProps = ctxProps
	return b
}

// SetDataConverter sets the data converter for the builder
func (b *WorkflowClientBuilder) SetDataConverter(dataConverter encoded.DataConverter) *WorkflowClientBuilder {
	b.dataConverter = dataConverter
	return b
}

// SetTracer sets the tracer for the builder
func (b *WorkflowClientBuilder) SetTracer(tracer opentracing.Tracer) *WorkflowClientBuilder {
	b.tracer = tracer
	return b
}

// BuildCadenceClient builds a client to cadence service
func (b *WorkflowClientBuilder) BuildCadenceClient() (client.Client, error) {
	service, err := b.BuildServiceClient()
	if err != nil {
		return nil, err
	}

	return client.NewClient(
		service,
		b.domain,
		&client.Options{
			Identity:           b.clientIdentity,
			MetricsScope:       b.metricsScope,
			DataConverter:      b.dataConverter,
			ContextPropagators: b.ctxProps,
			Tracer:             b.tracer,
			FeatureFlags: client.FeatureFlags{
				WorkflowExecutionAlreadyCompletedErrorEnabled: true,
			},
		}), nil
}

// BuildCadenceDomainClient builds a domain client to cadence service
func (b *WorkflowClientBuilder) BuildCadenceDomainClient() (client.DomainClient, error) {
	service, err := b.BuildServiceClient()
	if err != nil {
		return nil, err
	}

	return client.NewDomainClient(
		service,
		&client.Options{
			Identity:           b.clientIdentity,
			MetricsScope:       b.metricsScope,
			ContextPropagators: b.ctxProps,
			FeatureFlags: client.FeatureFlags{
				WorkflowExecutionAlreadyCompletedErrorEnabled: true,
			},
		},
	), nil
}

// BuildServiceClient builds a rpc service client to cadence service
func (b *WorkflowClientBuilder) BuildServiceClient() (workflowserviceclient.Interface, error) {
	if err := b.build(); err != nil {
		return nil, err
	}

	if b.dispatcher == nil {
		b.Logger.Fatal("No RPC dispatcher provided to create a connection to Cadence Service")
	}

	clientConfig := b.dispatcher.ClientConfig(_cadenceFrontendService)
	return compatibility.NewThrift2ProtoAdapter(
		apiv1.NewDomainAPIYARPCClient(clientConfig),
		apiv1.NewWorkflowAPIYARPCClient(clientConfig),
		apiv1.NewWorkerAPIYARPCClient(clientConfig),
		apiv1.NewVisibilityAPIYARPCClient(clientConfig),
	), nil
}

func (b *WorkflowClientBuilder) build() error {
	if b.dispatcher != nil {
		return nil
	}

	if len(b.hostPort) == 0 {
		return errors.New("HostPort is empty")
	}

	b.Logger.Debug("Creating RPC dispatcher outbound",
		zap.String("ServiceName", _cadenceFrontendService),
		zap.String("HostPort", b.hostPort))

	b.dispatcher = yarpc.NewDispatcher(yarpc.Config{
		Name: _cadenceClientName,
		Outbounds: yarpc.Outbounds{
			_cadenceFrontendService: {Unary: grpc.NewTransport().NewSingleOutbound(b.hostPort)},
		},
	})

	if b.dispatcher != nil {
		if err := b.dispatcher.Start(); err != nil {
			b.Logger.Fatal("Failed to create outbound transport channel: %v", zap.Error(err))
		}
	}
	return nil
}
