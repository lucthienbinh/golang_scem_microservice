package transport

import (
	"context"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/endpoint"
	pb "github.com/lucthienbinh/golang_scem_microservice/state_scem/pb"
	Repo "github.com/lucthienbinh/golang_scem_microservice/state_scem/repo"
)

type gRPCServer struct {
	// server is used to implement pb.StateScemServiceServer.
	pb.UnimplementedStateScemServiceServer
	deployWorkflow gt.Handler
	// createWorkflowInstance gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoint endpoint.Endpoints, logger log.Logger) pb.StateScemServiceServer {
	return &gRPCServer{
		deployWorkflow: gt.NewServer(
			endpoint.DeployWorkflowEndpoint,
			decodeDeployWorkflowRequest,
			encodeDeployWorkflowlResponse,
		),
		// createWorkflowInstance: gt.NewServer(
		// 	endpoint.CreateWorkflowInstance,
		// 	decodeMathRequest,
		// 	encodeMathResponse,
		// ),
	}
}

func (s *gRPCServer) DeployWorkflow(ctx context.Context, req *pb.DeployWorkflowRequest) (*pb.DeployWorkflowlResponse, error) {
	_, resp, err := s.deployWorkflow.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.DeployWorkflowlResponse), nil
}

func decodeDeployWorkflowRequest(_ context.Context, request interface{}) (interface{}, error) {
	workflowModelListDecoded := []Repo.WorkflowModel{}
	req := request.(*pb.DeployWorkflowRequest) // Struct from generated file
	for _, workflowModel := range req.WorkflowModel {
		workflowModelDecoded := Repo.WorkflowModel{}
		workflowModelDecoded.WorkflowProcessID = workflowModel.WorkflowProcessID
		workflowModelDecoded.WorkflowVersion = int(workflowModel.WorkflowVersion)
		workflowModelDecoded.WorkflowKey = workflowModel.WorkflowKey
		workflowModelDecoded.Step = int(workflowModel.Step)
		workflowModelDecoded.Type = int(workflowModel.Type)
		workflowModelDecoded.Name = workflowModel.Name
		workflowModelDecoded.NextStep1 = int(workflowModel.NextStep1)
		workflowModelDecoded.NextStep2 = int(workflowModel.NextStep2)
		workflowModelDecoded.ServiceRetry = int(workflowModel.ServiceRetry)
		workflowModelDecoded.MessageCorrelationName = workflowModel.MessageCorrelationName
		workflowModelListDecoded = append(workflowModelListDecoded)
	}
	return endpoint.DeployWorkflowRequest{WorkflowModelList: workflowModelListDecoded}, nil
}

func encodeDeployWorkflowlResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeployWorkflowlResponse)
	return &pb.DeployWorkflowlResponse{WorkflowKey: resp.WorkflowKey, Ok: resp.Ok}, nil // Struct rom generated file
}
