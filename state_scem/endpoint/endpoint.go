package endpoint

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/metrics"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/service"
)

// Endpoints structure
type Endpoints struct {
	DeployWorkflowEndpoint         endpoint.Endpoint
	CreateWorkflowInstanceEndpoint endpoint.Endpoint
}

type endpointCounter struct {
	service service.Service
	counter metrics.Counter
	latency metrics.Histogram
	logger  log.Logger
}

// MakeEndpoints function
func MakeEndpoints(_service service.Service, _logger log.Logger, _counter metrics.Counter, _latency metrics.Histogram) Endpoints {
	ec := endpointCounter{service: _service, counter: _counter, latency: _latency, logger: _logger}
	return Endpoints{
		DeployWorkflowEndpoint:         ec.makeDeployWorkflowEndpoint(),
		CreateWorkflowInstanceEndpoint: ec.makeCreateWorkflowInstanceEndpoint(),
	}
}

func (ec *endpointCounter) makeDeployWorkflowEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeployWorkflowRequest)
		workflowKey, ok, err := ec.service.DeployWorkflowService(ctx, req.WorkflowModelList)
		methodName := "Deployworkflow Endpoint"
		// COPY THIS STACK TO OBSERVE TIME
		defer func(begin time.Time) {
			lvs := []string{"method", methodName, "error", fmt.Sprint(err != nil)}
			ec.counter.With(lvs...).Add(1)
			ec.latency.With(lvs...).Observe(time.Since(begin).Seconds())
			level.Info(ec.logger).Log("method", methodName, "duration", time.Since(begin))
		}(time.Now())
		// COPY THIS STACK TO OBSERVE TIME
		return DeployWorkflowlResponse{WorkflowKey: workflowKey, Ok: ok}, err
	}
}

func (ec *endpointCounter) makeCreateWorkflowInstanceEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateWorkflowInstanceRequest)
		newInstanceID, ok, err := ec.service.CreateWorkflowInstanceService(ctx, req.WorkflowProcessID, req.WorkflowVariableList)
		methodName := "CreateWorkflowInstance Endpoint"
		// COPY THIS STACK TO OBSERVE TIME
		defer func(begin time.Time) {
			lvs := []string{"method", methodName, "error", fmt.Sprint(err != nil)}
			ec.counter.With(lvs...).Add(1)
			ec.latency.With(lvs...).Observe(time.Since(begin).Seconds())
			level.Info(ec.logger).Log("method", methodName, "duration", time.Since(begin))
		}(time.Now())
		// COPY THIS STACK TO OBSERVE TIME
		return CreateWorkflowInstanceResponse{WorkflowInstanceID: newInstanceID, Ok: ok}, err
	}
}
