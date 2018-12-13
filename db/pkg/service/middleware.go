package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(DbService) DbService

type loggingMiddleware struct {
	logger log.Logger
	next   DbService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a DbService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next DbService) DbService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Connect(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Connect", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Connect(ctx, s)
}
