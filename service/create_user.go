package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/fzlrhmn/simple-jwt/repository"
	"github.com/fzlrhmn/simple-jwt/util/crypt"
)

// CreateUser for create an user
func (us *UserSvc) CreateUser(ctx context.Context, req User) (*User, error) {
	pass, err := crypt.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := repository.User{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: pass,
	}

	_, err = us.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &req, nil
}
