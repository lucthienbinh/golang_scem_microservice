package endpoint

import (
	Repo "github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
)

type (
	// DeployWorkflowRequest structure
	DeployWorkflowRequest struct {
		WorkflowModelList []Repo.WorkflowModel `json:"workflow_model_list"`
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
		WorkflowProcessID    string                `json:"workflow_process_id"`
		WorkflowVariableList Repo.WorkflowVariable `json:"workflow_variable_list"`
	}
	// CreateWorkflowInstanceResponse structure
	CreateWorkflowInstanceResponse struct {
		WorkflowKey        string
		WorkflowInstanceID uint
		Ok                 bool
	}
)
