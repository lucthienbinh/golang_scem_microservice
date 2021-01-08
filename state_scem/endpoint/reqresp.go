package endpoint

import (
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
)

type (

	// CreateWorkflowInstanceRequest structure
	CreateWorkflowInstanceRequest struct {
		WorkflowKey      string `json:"workflow_key"`
		WorkflowVersion  int    `json:"workflow_version"`
		WorkflowVariable []service.WorkflowVariable
	}
	// CreateWorkflowInstanceResponse structure
	CreateWorkflowInstanceResponse struct {
		WorkflowInstanceID uint   `json:"workflow_instance_id"`
		Ok                 bool   `json:"ok"`
		Error              string `json:"error"`
	}
)
