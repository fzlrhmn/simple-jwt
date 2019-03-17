package endpoint

import (
	"context"

	"github.com/fzlrhmn/simple-jwt/service"
	e "github.com/fzlrhmn/simple-jwt/util/error"
	"github.com/go-kit/kit/endpoint"
)

// CreateUserRequest for holding create user request payload
type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUserResponse for user response structure
type CreateUserResponse struct {
	Username string `json:"username"`
}

// MakeCreateUserEndpoint return endpoint for create user
func MakeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(CreateUserRequest)
		if !ok {
			return nil, e.Unexpected()
		}

		u := service.User{
			Username: req.Username,
			Password: req.Password,
		}

		user, err := svc.CreateUser(ctx, u)
		if err != nil {
			return nil, err
		}

		return CreateUserResponse{
			Username: user.Username,
		}, nil
	}
}
