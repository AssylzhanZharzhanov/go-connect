package config

type AppConfig struct {
	Port string `env:"port" validate:"required"`
}
