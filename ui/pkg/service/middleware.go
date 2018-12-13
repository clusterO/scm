package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UiService) UiService

type loggingMiddleware struct {
	logger log.Logger
	next   UiService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UiService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UiService) UiService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Display(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Display", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Display(ctx, s)
}
