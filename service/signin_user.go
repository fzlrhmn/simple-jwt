package service

import (
	"context"
	"fmt"
	"time"

	"github.com/fzlrhmn/simple-jwt/util/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/fzlrhmn/simple-jwt/util/crypt"
	"github.com/spf13/viper"
)

// Claims is a struct for hold Claim
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SigninUser for signing in an user
func (us *UserSvc) SigninUser(ctx context.Context, req User) (*User, error) {
	u, err := us.Repository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	ok := crypt.CheckPasswordHash(req.Password, u.Password)
	if !ok {
		return nil, fmt.Errorf("Password doesn't match")
	}

	partnerExpiry := viper.GetDuration("app.expiry")
	expiresAt := time.Now().Add(partnerExpiry).Unix()

	stdClaim := jwt.StandardClaims{
		Issuer:    "Faisal Rahman",
		ExpiresAt: expiresAt,
	}

	claim := Claims{
		u.Username,
		stdClaim,
	}

	secret := config.Instance.GetString("app.secret")
	token, err := crypt.CreateToken(secret, claim)
	if err != nil {
		return nil, err
	}

	req.Password = ""
	req.Token = token

	return &req, nil
}
