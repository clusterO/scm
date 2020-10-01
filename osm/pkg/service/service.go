package service

import "context"

// OsmService describes the service.
type OsmService interface {
	// Add your methods here
	Locate(ctx context.Context, s string) (rs string, err error)
}

type basicOsmService struct{}

func (b *basicOsmService) Locate(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Locate
	return rs, err
}

// NewBasicOsmService returns a naive, stateless implementation of OsmService.
func NewBasicOsmService() OsmService {
	return &basicOsmService{}
}

// New returns a OsmService with all of the expected middleware wired in.
func New(middleware []Middleware) OsmService {
	var svc OsmService = NewBasicOsmService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
