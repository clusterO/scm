package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "scm/osm/pkg/service"
)

// LocateRequest collects the request parameters for the Locate method.
type LocateRequest struct {
	S string `json:"s"`
}

// LocateResponse collects the response parameters for the Locate method.
type LocateResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeLocateEndpoint returns an endpoint that invokes Locate on the service.
func MakeLocateEndpoint(s service.OsmService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LocateRequest)
		rs, err := s.Locate(ctx, req.S)
		return LocateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r LocateResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Locate implements Service. Primarily useful in a client.
func (e Endpoints) Locate(ctx context.Context, s string) (rs string, err error) {
	request := LocateRequest{S: s}
	response, err := e.LocateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LocateResponse).Rs, response.(LocateResponse).Err
}
