package crypt

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create new jwt token. This function has two inputs.
// issuer comes from "key" from kong / dynamodb.
// secret comes from "secret" from kong / dynamodb.
func CreateToken(secret string, c jwt.Claims) (string, error) {
	signingKey := []byte(secret)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// ParseToken is used for parse jwt token into object
func ParseToken(tokenString string, key string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
