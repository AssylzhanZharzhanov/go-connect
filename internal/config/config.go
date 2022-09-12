package config

type Config struct {
	Port string `environment:"port" binding:"required"`
}

func LoadConfig() error {
	return nil
}
