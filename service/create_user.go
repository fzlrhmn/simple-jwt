package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/fzlrhmn/simple-jwt/repository"
)

// CreateUser for create an user
func (us *UserSvc) CreateUser(ctx context.Context, req User) (*User, error) {
	user := repository.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: req.Password,
	}

	_, err := us.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &req, nil
}
