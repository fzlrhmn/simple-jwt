package service

import (
	"context"
	"testing"

	fake "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserAndSignin(t *testing.T) {
	ctx := context.Background()

	req := User{
		Username: fake.Username(),
		Password: fake.Password(true, true, true, false, false, 15),
	}

	user, err := svc.CreateUser(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	u, err := svc.SigninUser(ctx, req)
	assert.NoError(t, err)
	assert.NotEmpty(t, u)
}
