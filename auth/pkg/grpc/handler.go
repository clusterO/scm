package grpc

import (
	"context"
	endpoint "scm/auth/pkg/endpoint"
	pb "scm/auth/pkg/grpc/pb"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeAuthHandler creates the handler logic
func makeAuthHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AuthEndpoint, decodeAuthRequest, encodeAuthResponse, options...)
}

// decodeAuthResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeAuthRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.AuthRequest)
	return endpoint.AuthRequest{Email: req.Email, Content: req.Content}, nil
}

// encodeAuthResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAuthResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.AuthResponse)
	return &pb.AuthReply{Id: reply.Id}, nil
}
func (g *grpcServer) Auth(ctx context1.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	_, rep, err := g.auth.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AuthReply), nil
}
