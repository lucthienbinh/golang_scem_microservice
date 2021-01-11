package service

import (
	"context"

	Repo "github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
)

// Service interface
type Service interface {
	DeployWorkflow(ctx context.Context, workflowModelList []Repo.WorkflowModel) (string, bool, error)
	CreateWorkflowInstance(ctx context.Context, processID string, workflowVariableList []Repo.WorkflowVariable) (uint, bool, error)
}
