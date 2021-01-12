package endpoint

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
)

// Endpoints structure
type Endpoints struct {
	DeployWorkflowEndpoint         endpoint.Endpoint
	CreateWorkflowInstanceEndpoint endpoint.Endpoint
}

type endpointCounter struct {
	Service service.Service
	Counter metrics.Counter
	Latency metrics.Histogram
}

// MakeEndpoints function
func MakeEndpoints(s service.Service, counter metrics.Counter, latency metrics.Histogram) Endpoints {
	ec := endpointCounter{Service: s, Counter: counter, Latency: latency}
	return Endpoints{
		DeployWorkflowEndpoint:         ec.makeDeployWorkflowEndpoint(),
		CreateWorkflowInstanceEndpoint: makeCreateWorkflowInstanceEndpoint(s),
	}
}

func (ec *endpointCounter) makeDeployWorkflowEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeployWorkflowRequest)
		workflowKey, ok, err := ec.Service.DeployWorkflowService(ctx, req.WorkflowModelList)
		// COPY THIS STACK TO OBSERVE TIME
		defer func(begin time.Time) {
			lvs := []string{"method", "Deployworkflow Endpoint", "error", fmt.Sprint(err != nil)}
			ec.Counter.With(lvs...).Add(1)
			ec.Latency.With(lvs...).Observe(float64(time.Since(begin).Microseconds()))
		}(time.Now())
		// COPY THIS STACK TO OBSERVE TIME
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
