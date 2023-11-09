package common

import (
	"context"
	apiv1 "github.com/uber/cadence-idl/go/proto/api/v1"
	"go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/compatibility"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/transport/grpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type CadenceConfig struct {
	CadenceClient workflowserviceclient.Interface
	Logger        *zap.Logger
	Domain        string
	TaskList      string
}

var LOCadenceConfig *CadenceConfig

func NewCadenceConfig() *CadenceConfig {
	LOCadenceConfig = &CadenceConfig{}
	LOCadenceConfig.Setup()
	return LOCadenceConfig
}
func (c *CadenceConfig) Setup() {
	c.Logger = BuildLogger()
	c.CadenceClient = BuildCadenceClient()
	c.Domain = Domain
	c.TaskList = TaskListName
}

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

func (c *CadenceConfig) StartLoanOriginationWorkflowExecution(workflowType, workflowID, runID string, input []byte) {
	executionTimeout := int32(60)
	closeTimeout := int32(60)

	req := shared.StartWorkflowExecutionRequest{
		Domain:     &c.Domain,
		WorkflowId: &workflowID, // loanWorkflowID
		WorkflowType: &shared.WorkflowType{
			Name: &workflowType,
		},
		TaskList: &shared.TaskList{
			Name: &c.TaskList,
		},
		Input:                               input,
		ExecutionStartToCloseTimeoutSeconds: &executionTimeout,
		TaskStartToCloseTimeoutSeconds:      &closeTimeout,
		RequestId:                           &runID, //loanApplicationID
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resp, err := c.CadenceClient.StartWorkflowExecution(ctx, &req)
	if err != nil {
		c.Logger.Error("Failed to create workflow", zap.Error(err))
		panic("Failed to create workflow.")
	}

	c.Logger.Info("successfully started Loan Application workflow", zap.String("runID", resp.GetRunId()))
}
