package env

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Host string `envconfig:"HOST" default:"localhost"`
	Port int    `envconfig:"PORT" default:"8080"`
}

func LoadAppConfig() (AppConfig, error) {
	var config AppConfig
	if err := envconfig.Process("", &config); err != nil {
		return AppConfig{}, fmt.Errorf("failed to parse configuration from environment: %v", err)
	}

	return config, nil
}
