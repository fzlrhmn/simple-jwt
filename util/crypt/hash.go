package crypt

import (
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash ...
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// VerifyPassword ...
func VerifyPassword(s string) error {
	var (
		MinimumSeven, number, upper, special bool
		message                              string
	)

	letters := 0
	for _, s := range s {
		switch {
		case unicode.IsNumber(s):
			number = true
		case unicode.IsUpper(s):
			upper = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			special = true
		case unicode.IsLetter(s) || s == ' ':
		}
		letters++
	}
	MinimumSeven = letters >= 7

	message = ""
	if !MinimumSeven {
		message = message + "\nPassword must be seven character or more"
	}
	if !number {
		message = message + "\nPassword must contain minimum 1 number"
	}
	if !upper {
		message = message + "\nPassword must contain minimum 1 uppercasse"
	}
	if !special {
		message = message + "\nPassword must contain minimum 1 special character"
	}

	if message != "" {
		return fmt.Errorf(message)
	}

	return nil
}
