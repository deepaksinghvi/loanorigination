package loan_workflow

import (
	"context"
	"fmt"
	"github.com/deepaksinghvi/loanorigination/common"
	"github.com/deepaksinghvi/loanorigination/dto"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"strings"
	"time"
)

type conditionAndAction struct {
	// condition is a function pointer to a local activity
	condition interface{}
	// action is a function pointer to a regular activity
	action interface{}
}

var checks = []conditionAndAction{
	//{checkConditionLoanApproved, LoanFundingActivity},
	//{checkConditionLoanRejected, LoanRejectionActivity},
}

func LoanOriginationWorkflow(ctx workflow.Context, input dto.LoanApplicationInputStep) error {
	loanApplicationOutput := dto.LoanApplicationOutputStep{}

	/*ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 5,
		StartToCloseTimeout:    time.Minute * 5,
	}*/
	lao := workflow.LocalActivityOptions{
		// use short timeout as local activity is execute as function locally.
		ScheduleToCloseTimeout: time.Second,
	}
	ctx = workflow.WithLocalActivityOptions(ctx, lao)

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 5,
		StartToCloseTimeout:    time.Minute * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	state := common.Submitted
	var content string
	err := workflow.SetQueryHandler(ctx, common.QueryNameLoWorkflowState, func(includeContent bool) (common.QueryResult, error) {
		result := common.QueryResult{State: state}
		if includeContent {
			result.Content = content
		}
		return result, nil
	})
	if err != nil {
		return err
	}
	logger.Info("LoanOriginationWorkflow started")
	/*
		Activities are invoked asynchronously through task lists.
		A task list is essentially a queue used to store an activity task until it is picked up by an available worker.
	*/
	err = workflow.ExecuteActivity(ctx, LoanApplicationActivity, input).Get(ctx, &loanApplicationOutput)
	if err != nil {
		logger.Error("LoanApplicationActivity failed", zap.Error(err))
		return err
	}

	creditDecisionInput := dto.CreditDecisionInputStep{
		ApplicationNo: loanApplicationOutput.ApplicationNo,
		AadhaarNumber: input.AadhaarNumber,
		PanNumber:     input.PanNumber,
		AccountNo:     loanApplicationOutput.AccountNo,
	}
	creditDecisionOutput := dto.CreditDecisionOutputStep{}
	err = workflow.ExecuteActivity(ctx, CreditDecisionInternalActivity, creditDecisionInput).Get(ctx, &creditDecisionOutput)
	if err != nil {
		logger.Error("CreditDecisionInternalActivity failed", zap.Error(err))
		return err
	}

	var signalVal string
	signalName := "loan-signal"
	signalChan := workflow.GetSignalChannel(ctx, signalName)

	s := workflow.NewSelector(ctx)
	s.AddReceive(signalChan, func(c workflow.Channel, more bool) {
		c.Receive(ctx, &signalVal)
		workflow.GetLogger(ctx).Info("Received signal!", zap.String("signal", signalName), zap.String("value", signalVal))
	})
	s.Select(ctx)
	content = signalVal
	/*var conditionMeet bool
	err = workflow.ExecuteLocalActivity(ctx, checkConditionLoanApproved, signalVal).Get(ctx, &conditionMeet)
	if err != nil {
		return err
	}*/
	loanFundingInput := dto.LoanFundingInputStep{
		ApplicationNo: creditDecisionOutput.ApplicationNo,
		AccountNo:     creditDecisionOutput.AccountNo,
	}
	loanFundingOutput := dto.LoanApplicationOutputStep{}

	//if conditionMeet {
	if signalVal == "APPROVED" {
		err = workflow.ExecuteActivity(ctx, LoanFundingActivity, loanFundingInput).Get(ctx, &loanFundingOutput)
		if err != nil {
			logger.Error("LoanFundingActivity failed", zap.Error(err))
			return err
		}
		state = common.Approved
	} else {
		err = workflow.ExecuteActivity(ctx, LoanRejectionActivity, loanFundingInput).Get(ctx, &loanFundingOutput)
		if err != nil {
			logger.Error("LoanRejectionActivity failed", zap.Error(err))
			return err
		}
		state = common.Rejected
	}

	/*if len(signalVal) > 0 && signalVal != "APPROVED" {
		//return errors.New("signalVal")
		err = workflow.ExecuteActivity(ctx, LoanRejectionActivity, loanFundingInput).Get(ctx, &loanFundingOutput)
		if err != nil {
			logger.Error("LoanRejectionActivity failed", zap.Error(err))
			return err
		}
	} else {
		err = workflow.ExecuteActivity(ctx, LoanFundingActivity, loanFundingInput).Get(ctx, &loanFundingOutput)
		if err != nil {
			logger.Error("LoanFundingActivity failed", zap.Error(err))
			return err
		}
	}*/
	state = common.Closed
	logger.Info("Workflow result", zap.String("LoadApplicationOutput", fmt.Sprintf("State: %s, -  %v", state, loanFundingOutput)))
	return nil
}

func LoanOriginationIntegrationWorkflow(ctx workflow.Context, input dto.LoanApplicationInputStep) error {
	loanApplicationOutput := dto.LoanApplicationOutputStep{}

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("LoanOriginationIntegrationWorkflow started")

	err := workflow.ExecuteActivity(ctx, LoanApplicationActivity, input).Get(ctx, &loanApplicationOutput)
	if err != nil {
		logger.Error("LoanApplicationActivity failed", zap.Error(err))
		return err
	}
	creditDecisionInput := dto.CreditDecisionInputStep{
		ApplicationNo: loanApplicationOutput.ApplicationNo,
		AadhaarNumber: input.AadhaarNumber,
		PanNumber:     input.PanNumber,
		AccountNo:     loanApplicationOutput.AccountNo,
	}
	creditDecisionOutput := dto.CreditDecisionOutputStep{}
	err = workflow.ExecuteActivity(ctx, CreditDecisionExternalActivity, creditDecisionInput).Get(ctx, &creditDecisionOutput)
	if err != nil {
		logger.Error("CreditDecisionExternalActivity failed", zap.Error(err))
		return err
	}

	loanFundingInput := dto.LoanFundingInputStep{
		ApplicationNo: creditDecisionOutput.ApplicationNo,
		AccountNo:     creditDecisionOutput.AccountNo,
	}
	loanFundingOutput := dto.LoanApplicationOutputStep{}
	err = workflow.ExecuteActivity(ctx, LoanFundingActivity, loanFundingInput).Get(ctx, &loanFundingOutput)
	if err != nil {
		logger.Error("LoanFundingActivity failed", zap.Error(err))
		return err
	}
	logger.Info("Workflow result", zap.String("outData", fmt.Sprint("%v", loanFundingOutput)))
	return nil
}

func checkConditionLoanApproved(ctx context.Context, signal string) (bool, error) {
	// some real logic happen here...
	return strings.Contains(signal, "APPROVED"), nil
}

func checkConditionLoanRejected(ctx context.Context, signal string) (bool, error) {
	// some real logic happen here...
	return strings.Contains(signal, "REJECTED"), nil
}
