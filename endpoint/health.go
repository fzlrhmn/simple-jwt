package endpoint

import (
	"context"

	"github.com/fzlrhmn/simple-jwt/service"

	"github.com/go-kit/kit/endpoint"
)

// MakeHealthCheckEndpoint returns endpoint for health check
func MakeHealthCheckEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.HealthCheck(), nil
	}
}
