package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	req := User{
		Username: "Faisal",
		Password: "password123",
	}

	user, err := svc.CreateUser(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
}
