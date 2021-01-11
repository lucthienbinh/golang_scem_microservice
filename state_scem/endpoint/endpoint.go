package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
)

// Endpoints structure
type Endpoints struct {
	DeployWorkflowEndpoint         endpoint.Endpoint
	CreateWorkflowInstanceEndpoint endpoint.Endpoint
}

// MakeEndpoints function
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		DeployWorkflowEndpoint:         makeDeployWorkflowEndpoint(s),
		CreateWorkflowInstanceEndpoint: makeCreateWorkflowInstanceEndpoint(s),
	}
}

func makeDeployWorkflowEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeployWorkflowRequest)
		workflowKey, ok, err := s.DeployWorkflowService(ctx, req.WorkflowModelList)
		return DeployWorkflowlResponse{WorkflowKey: workflowKey, Ok: ok}, err
	}
}

func makeCreateWorkflowInstanceEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateWorkflowInstanceRequest)
		newInstanceID, ok, err := s.CreateWorkflowInstanceService(ctx, req.WorkflowProcessID, req.WorkflowVariableList)
		return CreateWorkflowInstanceResponse{WorkflowInstanceID: newInstanceID, Ok: ok}, err
	}
}
