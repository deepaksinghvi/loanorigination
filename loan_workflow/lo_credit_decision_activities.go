package loan_workflow

import (
	"context"
	"fmt"
	"github.com/deepaksinghvi/loanorigination/dto"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/activity"
)

func CreditDecisionInternalActivity(ctx context.Context, input dto.CreditDecisionInputStep) (dto.CreditDecisionOutputStep, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CreditDecisionInternalActivity started")
	output := dto.CreditDecisionOutputStep{
		ApplicationNo: input.ApplicationNo,
		AccountNo:     uuid.New(),
		CreditScore:   800.0,
	}
	logger.Info(fmt.Sprintf("CreditDecisionInternalActivity for application no %s! with credit score %f Completed.", input.ApplicationNo, output.CreditScore))
	return output, nil
}

func CreditDecisionExternalActivity(ctx context.Context, input dto.CreditDecisionInputStep) (dto.CreditDecisionOutputStep, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("CreditDecisionInternalActivity started")
	output := dto.CreditDecisionOutputStep{
		ApplicationNo: input.ApplicationNo,
		AccountNo:     uuid.New(),
		CreditScore:   800.0,
	}
	logger.Info(fmt.Sprintf("CreditDecisionInternalActivity for application no %s! with credit score %f Completed.", input.ApplicationNo, output.CreditScore))
	return output, nil
}

//
