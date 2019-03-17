package crypt

import (
	"flag"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	fake "github.com/brianvoe/gofakeit"
	jwt "github.com/dgrijalva/jwt-go"
	stringUtil "github.com/hooqtv/stairway/util/string"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	// Command line flag
	configFile := flag.String("config", "development", "configuration path")
	flag.Parse()

	path, _ := filepath.Abs(fmt.Sprintf("../../config/%s.toml", *configFile))
	fmt.Println("Loading configuration from: ", path)

	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed set find configuration: %s", err))
	}
}

type Claims struct {
	ChannelPartnerID string
	PartnerID        string
	jwt.StandardClaims
}

func TestCreateToken(t *testing.T) {
	issuer := viper.GetString("credentials.jwt_iss")
	fake.Seed(time.Now().UnixNano())
	key := stringUtil.String(32)
	ChannelPartnerID := "HOOQIND"
	partnerID := fake.BuzzWord()

	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	claim := jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: expiresAt,
	}

	claims := Claims{ChannelPartnerID, partnerID, claim}

	token, err := CreateToken(key, claims)
	assert.Nil(t, err)

	payload, err := ParseToken(token, key)

	assert.Nil(t, err)
	assert.Nil(t, payload.Valid())
	assert.Equal(t, partnerID, payload["PartnerID"])
	assert.Equal(t, ChannelPartnerID, payload["ChannelPartnerID"])
}
