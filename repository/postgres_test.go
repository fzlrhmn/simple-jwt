package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	fake "github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	user := User{
		ID:       uuid.New().String(),
		Username: fake.Username(),
		Password: fake.Password(true, true, true, false, false, 10),
	}

	u, err := repo.CreateUser(ctx, user)
	assert.NoError(t, err)
	assert.NotEmpty(t, u)
}
