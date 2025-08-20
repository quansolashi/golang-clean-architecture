package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	// Server config
	Port int64 `envconfig:"PORT" default:"8080"`

	// Database config
	DBSocket   string `envconfig:"DB_SOCKET" default:"tcp"`
	DBHost     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort     string `envconfig:"DB_PORT" default:"3306"`
	DBDatabase string `envconfig:"DB_DATABASE" default:""`
	DBUsername string `envconfig:"DB_USERNAME" default:"root"`
	DBPassword string `envconfig:"DB_PASSWORD" default:""`

	// JWT config
	JWTSecret string `envconfig:"JWT_SECRET" default:"secret"`

	// Log config
	LogLevel  string `envconfig:"LOG_LEVEL" default:"info"`
	LogOutput string `envconfig:"LOG_OUTPUT" default:"console"`
}

type Client interface {
	Load() (*Environment, error)
}

type client struct{}

func NewClient() Client {
	return &client{}
}

func (c *client) Load() (*Environment, error) {
	env := &Environment{}
	if err := processEnv("", env); err != nil {
		return nil, fmt.Errorf("failed to process environment: %w", err)
	}
	return env, nil
}

func processEnv(prefix string, spec interface{}) error {
	return envconfig.Process(prefix, spec)
}
