package config

// AppConfig - represents application configuration
type AppConfig struct {
	Port        string `env:"PORT" validate:"required"`
	DSN         string `env:"DSN" validate:"required"`
	ServiceName string `env:"SERVICE_NAME" validate:"required"`
	HostPort    string `env:"HOST_PORT" validate:"required"`
	Enable      bool   `env:"ENABLE" validate:"required"`
	LogSpans    bool   `env:"LOG_SPANS" validate:"required"`
}
