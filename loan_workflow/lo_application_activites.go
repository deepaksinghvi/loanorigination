package loan_workflow

import (
	"context"
	"fmt"
	"github.com/deepaksinghvi/loanorigination/dto"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/activity"
)

func LoanApplicationActivity(ctx context.Context, input dto.LoanApplicationInputStep) (dto.LoanApplicationOutputStep, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("LoanApplication started")
	output := dto.LoanApplicationOutputStep{
		ApplicationNo: activity.GetInfo(ctx).WorkflowExecution.RunID,
		AccountNo:     uuid.New(),
	}
	logger.Info(fmt.Sprintf("LoanApplication for application no %s! for account no %s Completed.", output.ApplicationNo, output.AccountNo))
	return output, nil
}
