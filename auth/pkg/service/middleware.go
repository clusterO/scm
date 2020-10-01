package service

import (
	"context"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(AuthService) AuthService

type loggingMiddleware struct {
	logger log.Logger
	next   AuthService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a AuthService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next AuthService) AuthService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Auth(ctx context.Context, email string, content string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Auth", "email", email, "content", content, "err", err)
	}()
	return l.next.Auth(ctx, email, content)
}
