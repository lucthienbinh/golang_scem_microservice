package endpoint

import (
	Repo "github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
)

type (
	// DeployWorkflowRequest structure
	DeployWorkflowRequest struct {
		WorkflowKey      string `json:"workflow_key"`
		WorkflowVersion  int    `json:"workflow_version"`
		WorkflowVariable []Repo.WorkflowVariable
	}
	// DeployWorkflowlResponse structure
	DeployWorkflowlResponse struct {
		Ok    bool   `json:"ok"`
		Error string `json:"error"`
	}
)

type (
	// CreateWorkflowInstanceRequest structure
	CreateWorkflowInstanceRequest struct {
		WorkflowKey      string `json:"workflow_key"`
		WorkflowVersion  int    `json:"workflow_version"`
		WorkflowVariable []Repo.WorkflowVariable
	}
	// CreateWorkflowInstanceResponse structure
	CreateWorkflowInstanceResponse struct {
		WorkflowInstanceID uint   `json:"workflow_instance_id"`
		Ok                 bool   `json:"ok"`
		Error              string `json:"error"`
	}
)
