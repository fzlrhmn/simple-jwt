package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	api "github.com/fzlrhmn/simple-jwt/endpoint"
	"github.com/fzlrhmn/simple-jwt/service"
	"github.com/fzlrhmn/simple-jwt/transport"
	utilpostgres "github.com/fzlrhmn/simple-jwt/util/connection/postgres"
	"github.com/fzlrhmn/simple-jwt/util/logger"
	"github.com/spf13/viper"
)

var (
	host *string
	port *int
)

func init() {
	// Command line flag
	port = flag.Int("port", 8000, "port")
	host = flag.String("host", "0.0.0.0", "host")
	configFile := flag.String("config", "config/development.toml", "configuration path")
	flag.Parse()

	viper.SetConfigFile(*configFile)

	// This will overwrite default config
	if configFromEnv := os.Getenv("USR_ENV"); configFromEnv != "" {
		viper.SetConfigFile(fmt.Sprintf("config/%s.toml", configFromEnv))
	}

	// This will overwrite the full config file path.
	if configFileFromEnv := os.Getenv("UMS_CONFIG_FILE"); configFileFromEnv != "" {
		viper.SetConfigFile(fmt.Sprintf("%s", configFileFromEnv))
	}

	// Bind environment variables to viper config
	viper.BindEnv("app.env", "USR_ENV")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if p := viper.GetInt("app.port"); p != 0 {
		port = &p
	}
}

func main() {
	errChan := make(chan error)
	env := viper.GetString("app.env")
	level := viper.GetString("log.level")
	logger, err := initLogger(env, level)
	if err != nil {
		return
	}

	logger.Info(fmt.Sprintf("Enviroment: %s", env))
	logger.Info(fmt.Sprintf("HTTP url: http://%s:%d", *host, *port))

	// Prepare postgres Clients
	utilpostgres.Initialize()

	svc := service.New()
	endpoint := api.New(svc)

	h := transport.MakeHTTPHandler(endpoint)
	s := &http.Server{Addr: fmt.Sprintf("%s:%d", *host, *port), Handler: h}

	defer func() {
		utilpostgres.GetInstance().Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cancel()

		logger.Warn("Shutting down UMS Server")
		s.Shutdown(ctx)
	}()

	go func() {
		errChan <- s.ListenAndServe()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	msg := <-errChan
	logger.Error(msg.Error())
}

func initLogger(env, levelS string) (*zap.Logger, error) {
	level, ok := logger.TranslateLevel(levelS)
	if !ok {
		return nil, fmt.Errorf("Log level \"%s\" is not valid", levelS)
	}

	logger, err := logger.GenerateConfig(env, level).Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
