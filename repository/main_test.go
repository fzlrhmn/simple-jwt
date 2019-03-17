package repository

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	fake "github.com/brianvoe/gofakeit"
	postgre "github.com/fzlrhmn/simple-jwt/util/connection/postgres"
	"github.com/spf13/viper"
)

var repo PostgresService

func TestMain(m *testing.M) {
	loadConfig()
	seedGofakeit()

	postgre.Initialize()
	repo = NewPostgresRepository()
	code := m.Run()
	os.Exit(code)
}

func loadConfig() {
	// Command line flag
	configFile := flag.String("config", "../config/development.toml", "configuration path")
	flag.Parse()

	viper.SetConfigFile(*configFile)

	// This will overwrite default config
	configFromEnv := os.Getenv("USR_ENV")
	if len(configFromEnv) > 0 {
		viper.SetConfigFile(fmt.Sprintf("../config/%s.toml", configFromEnv))
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func seedGofakeit() {
	fake.Seed(time.Now().UnixNano())
}
