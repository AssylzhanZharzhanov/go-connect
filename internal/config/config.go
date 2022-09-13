package config

type AppConfig struct {
	Port string `env:"PORT" validate:"required"`
	DSN  string `env:"DSN" validate:"required"`
}
