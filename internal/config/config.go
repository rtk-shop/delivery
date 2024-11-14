package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

const (
	DevEnv  = "development"
	ProdEnv = "production"
)

type Config struct {
	IsDev         bool
	Environment   string `env:"ENVIRONMENT,required"`
	Port          string `env:"PORT,required"`
	NovaPoshtaKey string `env:"NOVA_POSHTA_KEY,required"`
	NovaPoshtaURL string `env:"NOVA_POSHTA_API_URL,required"`
	RedisDB       int    `env:"REDIS_DB,required"`
	RedisPort     string `env:"REDIS_PORT,required"`
}

func New() *Config {

	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatalln("failed to parse config", "error", err)
	}

	cfg.IsDev = cfg.Environment == DevEnv

	return &cfg
}
