package config

import (
	"go/build"
	"os"
	"path"
)

// GetConfigPath will return application's config folder path
func GetConfigPath() string {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}
		configPath = path.Join(gopath, "src", "github.com", "fzlrhmn", "simple-jwt", "config")
	}
	return configPath
}

// GetConfigFile will reteurn application's config file path
func GetConfigFile(filepath string) string {
	return path.Join(GetConfigPath(), filepath)
}
