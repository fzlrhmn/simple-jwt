package repository

import "context"

// Postgres ...
type postgres struct{}

// PostgresService ...
type PostgresService interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	CreateUser(ctx context.Context, user User) (*User, error)
}

// NewPostgresRepository ...
func NewPostgresRepository() PostgresService {
	return &postgres{}
}
