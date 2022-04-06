package config

import (
	"github.com/auth_service/internal/apiserver"
	"github.com/auth_service/internal/store"
)

type Config struct {
	Server *apiserver.Config
	Store  *store.Config
}

func NewConfig() *Config {
	return &Config{
		Server: apiserver.NewConfig(),
		Store:  store.NewConfig(),
	}
}
