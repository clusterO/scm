package endpoint

import (
	"context"
	service "scm/auth/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// AuthRequest collects the request parameters for the Auth method.
type AuthRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

// AuthResponse collects the response parameters for the Auth method.
type AuthResponse struct {
	Id  string `json:"id"`
	Err error  `json:"err"`
}

// MakeAuthEndpoint returns an endpoint that invokes Auth on the service.
func MakeAuthEndpoint(s service.AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthRequest)
		id, err := s.Auth(ctx, req.Email, req.Content)
		return AuthResponse{
			Id:  id,
			Err: err,
		}, nil
	}
}

// Failed implements Failer.
func (r AuthResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Auth implements Service. Primarily useful in a client.
func (e Endpoints) Auth(ctx context.Context, s string) (rs string, err error) {
	request := AuthRequest{Email: s}
	response, err := e.AuthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(AuthResponse).Id, response.(AuthResponse).Err
}
