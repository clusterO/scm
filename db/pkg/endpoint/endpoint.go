package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "scm/db/pkg/service"
)

// ConnectRequest collects the request parameters for the Connect method.
type ConnectRequest struct {
	S string `json:"s"`
}

// ConnectResponse collects the response parameters for the Connect method.
type ConnectResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeConnectEndpoint returns an endpoint that invokes Connect on the service.
func MakeConnectEndpoint(s service.DbService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConnectRequest)
		rs, err := s.Connect(ctx, req.S)
		return ConnectResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r ConnectResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Connect implements Service. Primarily useful in a client.
func (e Endpoints) Connect(ctx context.Context, s string) (rs string, err error) {
	request := ConnectRequest{S: s}
	response, err := e.ConnectEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(ConnectResponse).Rs, response.(ConnectResponse).Err
}
