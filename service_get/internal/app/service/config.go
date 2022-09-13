package service

import "github.com/xamust/petbot/service_get/internal/app/store"

type Config struct {
	LogLevel string `toml:"log_level"`
	PortgRPC string `toml:"port_grpc"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		LogLevel: "info", //default param
		Store:    store.NewConfig(),
	}
}
