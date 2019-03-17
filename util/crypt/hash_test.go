package crypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	// should generate hash
	hash, _ := HashPassword("devyo1")
	assert.NotEmpty(t, hash)

	// should return true
	res := CheckPasswordHash("devyo1", hash)
	assert.Equal(t, true, res)

	// should return false
	res = CheckPasswordHash("devyo1", "invalid-hash")
	assert.Equal(t, false, res)
}

func TestVerifyPassword(t *testing.T) {
	cases := map[string]struct {
		Hash           string
		Error          bool
		ExpectedResult string
	}{
		"LessThanSeven": {
			Hash:           "Less7*",
			Error:          true,
			ExpectedResult: "\nPassword must be seven character or more",
		},
		"NoNumber": {
			Hash:           "NoNumber&&*^",
			Error:          true,
			ExpectedResult: "\nPassword must contain minimum 1 number",
		},
		"NoUppercase": {
			Hash:           "alllowercase12*&",
			Error:          true,
			ExpectedResult: "\nPassword must contain minimum 1 uppercasse",
		},
		"NoSpecial": {
			Hash:           "NoSp3cial",
			Error:          true,
			ExpectedResult: "\nPassword must contain minimum 1 special character",
		},
		"AllPassed": {
			Hash:           "*AllT3stP4ss3D*",
			Error:          false,
			ExpectedResult: "",
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			err := VerifyPassword(test.Hash)
			if test.Error {
				assert.Error(t, err)
				assert.Equal(t, test.ExpectedResult, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
