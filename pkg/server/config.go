package server

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Environment Vars
const (
	EnvEnvironment    = "ENVIRONMENT"
	EnvConfigLocation = "CONFIG_PATH"
)

// Config exposes Server configuration options
type Config struct {
	BasePath       string
	MetricsEnabled bool
	Name           string
	Port           int
}

// NewConfig constructs a new *server.Config instance with defaults
func NewConfig() *Config {
	viper.SetConfigName(getEnvironmentVariable(EnvEnvironment, "review"))
	viper.AddConfigPath(getEnvironmentVariable(EnvConfigLocation, "./config"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	return &Config{
		BasePath:       "/api/v1",
		Name:           viper.GetString("name"),
		Port:           viper.GetInt("port"),
		MetricsEnabled: viper.GetBool("metrics"),
	}
}

func getEnvironmentVariable(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
