package common

const (
	HostPort = "127.0.0.1:7833"
	Domain   = "lo-domain"
	// TaskListName identifies set of loan_client workflows, activities, and workers.
	// It could be your group or loan_client or application name.
	TaskListName   = "lo-loan_worker"
	ClientName     = "lo-loan_worker"
	CadenceService = "cadence-frontend"

	QueryNameLoWorkflowState = "lo-workflow-state"

	Initialized State = "initialized"
	Submitted   State = "submitted"
	Approved    State = "approved"
	Rejected    State = "rejected"
	Closed      State = "closed"
)

type State string
type QueryResult struct {
	State   State
	Content string
}

type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}
