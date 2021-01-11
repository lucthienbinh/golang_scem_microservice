package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
)

// Endpoints structure
type Endpoints struct {
	DeployWorkflow         endpoint.Endpoint
	CreateWorkflowInstance endpoint.Endpoint
}

// MakeEndpoints function
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		DeployWorkflow:         makeDeployWorkflowEndpoint(s),
		CreateWorkflowInstance: makeCreateWorkflowInstanceEndpoint(s),
	}
}

func makeDeployWorkflowEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeployWorkflowRequest)
		workflowKey, ok, err := s.DeployWorkflow(ctx, req.WorkflowModelList)
		return DeployWorkflowlResponse{WorkflowKey: workflowKey, Ok: ok}, err
	}
}

func makeCreateWorkflowInstanceEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateWorkflowInstanceRequest)
		newInstanceID, ok, err := s.CreateWorkflowInstance(ctx, req.WorkflowProcessID, req.WorkflowVariableList)
		return CreateWorkflowInstanceResponse{WorkflowInstanceID: newInstanceID, Ok: ok}, err
	}
}
