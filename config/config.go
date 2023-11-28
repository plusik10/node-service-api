package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		GRPC `yaml:"grpc"`
		PG   `yaml:"postgres"`
		Log  `yaml:"logger"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	GRPC struct {
		Port string `yaml:"port"`
	}

	PG struct {
		URL string `yaml:"url" env:"PG_URL"`
	}
	Log struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return cfg, err
}
