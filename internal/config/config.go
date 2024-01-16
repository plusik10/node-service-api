package config

import (
	"errors"
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v4/pgxpool"
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
		DSN                string `yaml:"dsn" env:"PG_DSN"`
		MaxOpenConnections int32  `yaml:"max_connections"  env:"PG_MAX_CONNECT"`
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
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
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

func (c *Config) GetDBConfig() (*pgxpool.Config, error) {
	poolConfig, err := pgxpool.ParseConfig(c.PG.DSN)
	if err != nil {
		return nil, err
	}

	poolConfig.ConnConfig.BuildStatementCache = nil
	poolConfig.ConnConfig.PreferSimpleProtocol = true
	poolConfig.MaxConns = c.PG.MaxOpenConnections

	return poolConfig, nil
}
