package endpoint

import (
	"context"

	"github.com/fzlrhmn/simple-jwt/service"
	e "github.com/fzlrhmn/simple-jwt/util/error"
	"github.com/go-kit/kit/endpoint"
)

// SigninUserRequest for holding signin user request payload
type SigninUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SigninUserResponse for signin user response structure
type SigninUserResponse struct {
	Token string `json:"token"`
}

// MakeSigninUserEndpoint return endpoint for create user
func MakeSigninUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(SigninUserRequest)
		if !ok {
			return nil, e.Unexpected()
		}

		u := service.User{
			Username: req.Username,
			Password: req.Password,
		}

		user, err := svc.SigninUser(ctx, u)
		if err != nil {
			return nil, err
		}

		return SigninUserResponse{
			Token: user.Token,
		}, nil
	}
}
