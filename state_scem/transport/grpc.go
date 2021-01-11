package transport

import (
	"context"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/endpoint"
)

type gRPCServer struct {
	deployWorkflow gt.Handler
	// createWorkflowInstance gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoint endpoint.Endpoints, logger log.Logger) pb.MathqServiceServer {
	return &gRPCServer{
		deployWorkflow: gt.NewServer(
			endpoint.DeployWorkflow,
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

func (s *gRPCServer) DeployWorkflow(ctx context.Context, req *pb.MathqRequest) (*pb.MathqResponse, error) {
	_, resp, err := s.deployWorkflow.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathqResponse), nil
}

// func (s *gRPCServer) CreateWorkflowInstance(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
// 	_, resp, err := s.createWorkflowInstance.ServeGRPC(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.(*pb.MathResponse), nil
// }

func decodeDeployWorkflowRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest) // From generated file
	return endpoint.DeployWorkflowRequest{WorkflowModelList: req.NuqmA}, nil
}

func encodeDeployWorkflowlResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeployWorkflowlResponse)
	return &pb.MathResponse{Result: resp.Resqult}, nil // From generated file
}
