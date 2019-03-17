package service

import (
	"context"

	"github.com/fzlrhmn/simple-jwt/repository"
)

type (

	// UserService is an interface for main service
	UserService interface {
		HealthCheck() bool
		CreateUser(context.Context, User) (*User, error)
	}

	// UserSvc is representation of User service instance
	UserSvc struct {
		Repository repository.PostgresService
	}

	// User is a struct to holds user data
	User struct {
		Username string
		Password string
		JWT      string
	}
)

// New is creating new service instance
func New() UserService {
	return &UserSvc{
		Repository: repository.NewPostgresRepository(),
	}
}
