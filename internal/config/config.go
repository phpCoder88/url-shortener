package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port            uint          `default:"8000"`
	IdleTimeout     time.Duration `default:"2m"`
	ReadTimeout     time.Duration `default:"5s"`
	WriteTimeout    time.Duration `default:"5s"`
	ShutdownTimeout time.Duration `default:"15s"`
}

func GetConfig() (*Config, error) {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
