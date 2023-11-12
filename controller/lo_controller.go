package controller

import (
	"context"
	"encoding/json"
	"github.com/deepaksinghvi/loanorigination/common"
	"github.com/deepaksinghvi/loanorigination/dto"
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
	"net/http"
)

// CreateLoanApplication godoc
// @Summary Create New Loan APplication
// @ID create-loan-application
// @Tags loan-origination
// @Accept json
// @Produce json
// @Param loanApplicationRequest body dto.LoanApplicationInputStep true "Loan Application "
// @Success 200 {object} string
// @Failure 500 {object} common.HTTPError
// @Router /loan-application [post]
func CreateLoanApplication(c *gin.Context) {
	// Validate input
	var loanApplicationInput dto.LoanApplicationInputStep
	// Call BindJSON to bind the received JSON to input.
	if err := c.ShouldBindJSON(&loanApplicationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	workflowID := uuid.New()
	loanApplicationID := uuid.New()
	input, _ := json.Marshal(loanApplicationInput)
	common.LOCadenceHelper.StartLoanOriginationWorkflowExecution("loan_workflow.LoanOriginationWorkflow", workflowID, loanApplicationID, input)
	c.JSON(http.StatusOK, gin.H{"data": loanApplicationID})
}

// GetLoanApplication godoc
// @Summary GetLoan Application state
// @ID get-loan-application
// @Tags loan-origination
// @Accept json
// @Produce json
// @Param workflow_id path string true "Workflow ID"
// @Param run_id path string true "Run ID"
// @Success 200 {object} common.QueryResult
// @Failure 500 {object} common.HTTPError
// @Router /loan-application/{workflow_id}/{run_id} [get]
func GetLoanApplication(c *gin.Context) {
	workflowID := c.Param("workflow_id")
	runID := c.Param("run_id")
	queryResult, err := queryLoanOriginationApplicationState(workflowID, runID, true)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": queryResult})
}

func queryLoanOriginationApplicationState(workflowID, runID string, args ...interface{}) (*common.QueryResult, error) {
	var result common.QueryResult
	workflowClient, err := common.LOCadenceHelper.Builder.BuildCadenceClient()
	resp, err := workflowClient.QueryWorkflowWithOptions(context.Background(),
		&client.QueryWorkflowWithOptionsRequest{
			WorkflowID:            workflowID,
			RunID:                 runID,
			QueryType:             common.QueryNameLoWorkflowState,
			QueryConsistencyLevel: shared.QueryConsistencyLevelStrong.Ptr(),
			Args:                  args,
		})
	if err != nil {
		common.LOCadenceHelper.Logger.Error("Failed to query workflow", zap.Error(err))
		//panic("Failed to query workflow.")
	}
	if err := resp.QueryResult.Get(&result); err != nil {
		common.LOCadenceHelper.Logger.Error("Failed to decode query result", zap.Error(err))
	}
	common.LOCadenceHelper.Logger.Info("Received consistent query result.", zap.Any("Result", result))
	return &result, err
}
