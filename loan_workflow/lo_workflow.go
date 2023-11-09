package loan_workflow

import (
	"errors"
	"fmt"
	"github.com/deepaksinghvi/loanorigination/dto"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"time"
)

func LoanOriginationWorkflow(ctx workflow.Context, input dto.LoanApplicationInputStep) error {
	loanApplicationOutput := dto.LoanApplicationOutputStep{}

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 5,
		StartToCloseTimeout:    time.Minute * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("LoanOriginationWorkflow started")

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

	if len(signalVal) > 0 && signalVal != "SOME_VALUE" {
		return errors.New("signalVal")
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
