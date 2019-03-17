package endpoint

import (
	"github.com/fzlrhmn/simple-jwt/service"
	"github.com/go-kit/kit/endpoint"
)

// Set is collection of gokit endpoints implementation
type Set struct {
	GetHealthCheckEndpoint endpoint.Endpoint
	CreateUserEndpoint     endpoint.Endpoint
}

// New is for create instance of endpoint
func New(svc service.UserService) Set {
	return Set{
		GetHealthCheckEndpoint: MakeHealthCheckEndpoint(svc),
		CreateUserEndpoint:     MakeCreateUserEndpoint(svc),
	}
}
