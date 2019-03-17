package repository

import (
	"context"

	utilPostgres "github.com/fzlrhmn/simple-jwt/util/connection/postgres"
)

func (*postgres) GetUserByUsername(ctx context.Context, username string) (*User, error) {
	pg := utilPostgres.GetInstance()

	user := new(User)

	err := pg.Model(user).Where("username = ?", username).First()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (*postgres) CreateUser(ctx context.Context, user User) (*User, error) {
	pg := utilPostgres.GetInstance()

	err := pg.Insert(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
