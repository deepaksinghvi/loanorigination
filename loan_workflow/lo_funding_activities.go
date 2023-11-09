package loan_workflow

import (
	"context"
	"fmt"
	"github.com/deepaksinghvi/loanorigination/dto"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/activity"
)

func LoanFundingActivity(ctx context.Context, input dto.LoanFundingInputStep) (dto.LoanFundingOutputStep, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("LoanFundingActivity started")
	output := dto.LoanFundingOutputStep{
		ApplicationNo:      input.ApplicationNo,
		AccountNo:          uuid.New(),
		DisbursementAmount: 1000.0,
	}
	logger.Info(fmt.Sprintf("LoanFundingActivity for application no %s! with disbursement amount %f Completed.", input.ApplicationNo, output.DisbursementAmount))
	return output, nil
}
