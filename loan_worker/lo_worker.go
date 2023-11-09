package loan_worker

import (
	"github.com/deepaksinghvi/loanorigination/common"
	"github.com/deepaksinghvi/loanorigination/loan_workflow"
	"github.com/uber-go/tally"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/worker"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// StartWorker creates and starts a basic Cadence loan_worker.
func StartWorker() {
	logger, cadenceClient := common.BuildLogger(), common.BuildCadenceClient()
	workerOptions := worker.Options{
		Logger:       logger,
		MetricsScope: tally.NewTestScope(common.TaskListName, nil),
	}

	w := worker.New(
		cadenceClient,
		common.Domain,
		common.TaskListName,
		workerOptions)
	// HelloWorld loan_workflow registration
	w.RegisterWorkflowWithOptions(loan_workflow.LoanOriginationWorkflow, workflow.RegisterOptions{Name: "loan_workflow.LoanOriginationWorkflow"})
	w.RegisterActivityWithOptions(loan_workflow.LoanApplicationActivity, activity.RegisterOptions{Name: "loan_workflow.LoanApplicationActivity"})
	w.RegisterActivityWithOptions(loan_workflow.CreditDecisionInternalActivity, activity.RegisterOptions{Name: "loan_workflow.CreditDecisionInternalActivity"})
	w.RegisterActivityWithOptions(loan_workflow.LoanFundingActivity, activity.RegisterOptions{Name: "loan_workflow.LoanFundingActivity"})

	err := w.Start()
	if err != nil {
		panic("Failed to start loan_worker: " + err.Error())
	}
	logger.Info("Started Worker.", zap.String("loan_worker", common.TaskListName))

}
