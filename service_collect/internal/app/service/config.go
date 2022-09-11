package service

import "github.com/xamust/petbot/service_collect/internal/app/store"

type Config struct {
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		LogLevel: "info", //default param
		Store:    store.NewConfig(),
	}
}
