package transport

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	api "github.com/fzlrhmn/simple-jwt/endpoint"
	"github.com/fzlrhmn/simple-jwt/service"
	utilpostgres "github.com/fzlrhmn/simple-jwt/util/connection/postgres"
	"github.com/spf13/viper"
)

var (
	server *httptest.Server
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	// Command line flag
	configFile := flag.String("config", "../config/development.toml", "configuration path")
	flag.Parse()

	viper.SetConfigFile(*configFile)

	// This will overwrite default config
	configFromEnv := os.Getenv("USR_ENV")
	if len(configFromEnv) > 0 {
		viper.SetConfigFile(fmt.Sprintf("../../config/ap-southeast-1/%s.toml", configFromEnv))
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Prepare postgres Clients
	utilpostgres.Initialize()

	svc := service.New()
	endpoint := api.New(svc)
	handler := MakeHTTPHandler(endpoint)
	server = httptest.NewServer(handler)

	// Fake seeder
	gofakeit.Seed(time.Now().UnixNano())
}

func unmarshalError(t *testing.T, body []byte) string {
	var e *ErrorResponse
	err := json.Unmarshal(body, &e)
	if err != nil {
		t.Fatalf("Failed unmarshaling error from response body: %s", err.Error())
	}

	return e.Error[0].Code
}

func checkError(t *testing.T, r *http.Response) string {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Failed reading response body: %s", err.Error())
	}
	defer r.Body.Close()

	return unmarshalError(t, resp)
}
