package transport

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/endpoint"
	pb "github.com/lucthienbinh/golang_scem_microservice/state_scem/pb"
	"google.golang.org/protobuf/encoding/protojson"
)

type gRPCServer struct {
	// server is used to implement pb.StateScemServiceServer.
	pb.UnimplementedStateScemServiceServer
	deployWorkflow         gt.Handler
	createWorkflowInstance gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoint endpoint.Endpoints, logger log.Logger) pb.StateScemServiceServer {
	return &gRPCServer{
		deployWorkflow: gt.NewServer(
			endpoint.DeployWorkflowEndpoint,
			decodeDeployWorkflowRequest,
			encodeDeployWorkflowlResponse,
		),
		createWorkflowInstance: gt.NewServer(
			endpoint.CreateWorkflowInstanceEndpoint,
			decodeCreateWorkflowInstanceRequest,
			encodeCreateWorkflowInstanceResponse,
		),
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
	req := request.(*pb.DeployWorkflowRequest)
	requestDecoded := endpoint.DeployWorkflowRequest{}
	// Parse to jsonpb of pb
	byteValue, _ := protojson.Marshal(req)
	// Parse json to struct
	json.Unmarshal(byteValue, &requestDecoded)
	return endpoint.DeployWorkflowRequest{WorkflowModelList: requestDecoded.WorkflowModelList}, nil
}

func encodeDeployWorkflowlResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.DeployWorkflowlResponse)
	return &pb.DeployWorkflowlResponse{WorkflowKey: resp.WorkflowKey, Ok: resp.Ok}, nil
}

func decodeCreateWorkflowInstanceRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateWorkflowInstanceRequest)
	requestDecoded := endpoint.CreateWorkflowInstanceRequest{}
	// Parse to jsonpb of pb
	byteValue, _ := protojson.Marshal(req)
	// Parse json to struct
	json.Unmarshal(byteValue, &requestDecoded)
	return endpoint.CreateWorkflowInstanceRequest{
		WorkflowProcessID:    requestDecoded.WorkflowProcessID,
		WorkflowVariableList: requestDecoded.WorkflowVariableList}, nil
}

func encodeCreateWorkflowInstanceResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateWorkflowInstanceResponse)
	return &pb.CreateWorkflowInstanceResponse{WorkflowInstanceID: int32(resp.WorkflowInstanceID), Ok: resp.Ok}, nil
}
