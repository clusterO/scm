package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(OsmService) OsmService

type loggingMiddleware struct {
	logger log.Logger
	next   OsmService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a OsmService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next OsmService) OsmService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Locate(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Locate", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Locate(ctx, s)
}
