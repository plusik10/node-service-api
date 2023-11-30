package config

import (
	"errors"
	"flag"
	"fmt"
	"os"

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
		URL string `yaml:"url" env:"PG_DSN"`
	}
	Log struct {
		Level string `yaml:"log_level" env:"LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	path := fetchConfigPath()
	if path == "" {
		return nil, errors.New("config path is empty")
	}

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

func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
