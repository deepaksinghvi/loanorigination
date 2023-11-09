package main

import (
	"context"
	"encoding/json"
	"github.com/deepaksinghvi/loanorigination/common"
	"github.com/deepaksinghvi/loanorigination/dto"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/zap"
	"time"
)

/**
A client program can also be used to execute workflow.
*/
func main() {
	logger, cadenceClient := common.BuildLogger(), common.BuildCadenceClient()
	domain := common.Domain
	tasklist := common.TaskListName
	loanWorkflowID := uuid.New()
	loanApplicationID := uuid.New()
	executionTimeout := int32(60)
	closeTimeout := int32(60)

	workflowType := "loan_workflow.LoanOriginationWorkflow"
	loanApplicationInput := dto.LoanApplicationInputStep{
		//ApplicationNo: uuid.New(),
		ApplicantName: "John Doe",
		AadhaarNumber: "123456781234",
		PanNumber:     "CKUKN1234T",
	}
	input, _ := json.Marshal(loanApplicationInput)
	req := shared.StartWorkflowExecutionRequest{
		Domain:     &domain,
		WorkflowId: &loanWorkflowID,
		WorkflowType: &shared.WorkflowType{
			Name: &workflowType,
		},
		TaskList: &shared.TaskList{
			Name: &tasklist,
		},
		Input:                               input,
		ExecutionStartToCloseTimeoutSeconds: &executionTimeout,
		TaskStartToCloseTimeoutSeconds:      &closeTimeout,
		RequestId:                           &loanApplicationID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resp, err := cadenceClient.StartWorkflowExecution(ctx, &req)
	if err != nil {
		logger.Error("Failed to create workflow", zap.Error(err))
		panic("Failed to create workflow.")
	}

	logger.Info("successfully started Loan Application workflow", zap.String("runID", resp.GetRunId()))
}
