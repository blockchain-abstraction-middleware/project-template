package config

import (
	"fmt"
	"os"

	"github.com/blockchain-abstraction-middleware/decrypter/decrypt"
	"github.com/ghodss/yaml"
)

// Environment Vars
const (
	EnvConfigName     = "CONFIG_NAME"
	EnvConfigLocation = "CONFIG_PATH"
)

// Config exposes service configuration options
type Config struct {
	Environment    string `yaml:"environment"`
	Name           string `yaml:"name"`
	Port           int    `yaml:"port"`
	MetricsEnabled bool   `yaml:"metrics"`
}

// NewConfig constructs a new *server.Config instance with defaults
func NewConfig() *Config {
	var config Config

	decryptedFile, err := decrypt.Decrypt(getEnvironmentVariable(EnvConfigLocation, "./config/"), getEnvironmentVariable(EnvConfigName, "review.yml"))
	if err != nil {
		panic(fmt.Errorf("Fatal error decrypting config file: %s", err))
	}

	err = yaml.Unmarshal(decryptedFile, &config)
	if err != nil {
		panic(fmt.Errorf("Fatal error parsing into config file: %s", err))
	}

	return &config
}

func getEnvironmentVariable(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
