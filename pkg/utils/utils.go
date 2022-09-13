package utils

import (
	"fmt"

	"github.com/AssylzhanZharzhanov/connect/internal/config"

	"github.com/caarlos0/env"
)

// LoadConfig - reads configuration from the environment variables.
func LoadConfig() (config config.AppConfig, _ error) {
	if err := env.Parse(&config); err != nil {
		return config, fmt.Errorf("failed to read configuration: %v", err)
	}
	return config, nil
}
