package service

import (
	"context"

	"github.com/satori/go.uuid"
)

// AuthService describes the service.
type AuthService interface {
	// Add your methods here
	Auth(ctx context.Context, email string, contnet string) (Id string, err error)
}

type basicAuthService struct{}

func (b *basicAuthService) Auth(ctx context.Context, email string, content string) (Id string, err error) {
	// TODO implement the business logic of Auth
	id := uuid.NewV4()
	return id.String(), nil
}

// NewBasicAuthService returns a naive, stateless implementation of AuthService.
func NewBasicAuthService() AuthService {
	return &basicAuthService{}
}

// New returns a AuthService with all of the expected middleware wired in.
func New(middleware []Middleware) AuthService {
	var svc AuthService = NewBasicAuthService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
