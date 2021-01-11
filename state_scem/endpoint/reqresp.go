package endpoint

import (
	Repo "github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
)

type (
	// DeployWorkflowRequest structure
	DeployWorkflowRequest struct {
		WorkflowModelList []Repo.WorkflowModel
	}
	// DeployWorkflowlResponse structure
	DeployWorkflowlResponse struct {
		WorkflowKey string
		Ok          bool
	}
)

type (
	// CreateWorkflowInstanceRequest structure
	CreateWorkflowInstanceRequest struct {
		WorkflowProcessID    string
		WorkflowVariableList []Repo.WorkflowVariable
	}
	// CreateWorkflowInstanceResponse structure
	CreateWorkflowInstanceResponse struct {
		WorkflowInstanceID uint
		Ok                 bool
	}
)
