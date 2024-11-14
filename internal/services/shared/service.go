package shared

import (
	"rtk/delivery/internal/config"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	PopularCities() ([]byte, error)
}

type service struct {
	config *config.Config
	cache  *redis.Client
}

func New(config *config.Config, cache *redis.Client) Service {
	return &service{
		config: config,
		cache:  cache,
	}
}
