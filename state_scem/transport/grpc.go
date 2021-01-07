package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/junereycasuga/gokit-grpc-demo/pb"
	"github.com/lucthienbinh/golang_scem_microservice/state_scem/endpoint"
)

type gRPCServer struct {
	add      gt.Handler
	subtract gt.Handler
	multiply gt.Handler
	divide   gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoint endpoint.Endpoints, logger log.Logger) pb.MathServiceServer {
	return &gRPCServer{
		add: gt.NewServer(
			endpoint.Add,
			decodeMathRequest,
			encodeMathResponse,
		),
		subtract: gt.NewServer(
			endpoint.Subtract,
			decodeMathRequest,
			encodeMathResponse,
		),
	}
}

func (s *gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Subtract(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.subtract.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest)
	return endpoint.MathReq{NumA: req.NumA, NumB: req.NumB}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.MathResp)
	return &pb.MathResponse{Result: resp.Result}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
