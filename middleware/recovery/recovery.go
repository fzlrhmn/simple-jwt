package recovery

import (
	"context"
	"errors"

	er "github.com/fzlrhmn/simple-jwt/util/error"
	"github.com/go-kit/kit/endpoint"
)

// CreateMiddleware creates an endpoint-layer middleware to recover from panic
func CreateMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (result interface{}, e error) {
			defer func() {
				if r := recover(); r != nil {
					msg, ok := r.(string)
					if !ok {
						err, ok := r.(error)
						if ok {
							msg = err.Error()
						} else {
							msg = "Unexpected panic"
						}
					}
					e = er.UnexpectedPanic().Wrap(errors.New(msg))
				}
			}()
			return next(ctx, request)
		}
	}
}
