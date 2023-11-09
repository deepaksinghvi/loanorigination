package controller

import (
	"encoding/json"
	"github.com/deepaksinghvi/loanorigination/common"
	"github.com/deepaksinghvi/loanorigination/dto"
	"github.com/gin-gonic/gin"
	"github.com/pborman/uuid"
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
	common.LOCadenceConfig.StartLoanOriginationWorkflowExecution("loan_workflow.LoanOriginationWorkflow", workflowID, loanApplicationID, input)
	c.JSON(http.StatusOK, gin.H{"data": loanApplicationID})
}
