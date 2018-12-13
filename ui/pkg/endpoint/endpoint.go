package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "scm/ui/pkg/service"
)

// DisplayRequest collects the request parameters for the Display method.
type DisplayRequest struct {
	S string `json:"s"`
}

// DisplayResponse collects the response parameters for the Display method.
type DisplayResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeDisplayEndpoint returns an endpoint that invokes Display on the service.
func MakeDisplayEndpoint(s service.UiService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DisplayRequest)
		rs, err := s.Display(ctx, req.S)
		return DisplayResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r DisplayResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Display implements Service. Primarily useful in a client.
func (e Endpoints) Display(ctx context.Context, s string) (rs string, err error) {
	request := DisplayRequest{S: s}
	response, err := e.DisplayEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DisplayResponse).Rs, response.(DisplayResponse).Err
}
