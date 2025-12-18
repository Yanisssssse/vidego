package server

import (
	"github.com/Yanisssssse/vidego/internal/storage"
)

type Config struct {
	Host         string
	Port         string
	AuthRequired bool
	Storage      storage.Storage
}

func DefaultConfig() (*Config, error) {
	s, err := storage.NewLocalStorage("./data")
	if err != nil {
		return nil, err
	}
	return &Config{
		Host:         "localhost",
		Port:         "7789",
		AuthRequired: false,
		Storage:      s,
	}, nil
}

func NewConfig(host string, port string, requireAuth bool, storage storage.Storage) *Config {
	return &Config{
		Host:         host,
		Port:         port,
		AuthRequired: requireAuth,
		Storage:      storage,
	}
}
