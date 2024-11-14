package novaposhta

import (
	"rtk/delivery/internal/config"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	Warehouses(cityID string) ([]byte, error)
}

type service struct {
	config *config.Config
	cache  *redis.Client
	apiKey string
}

func New(config *config.Config, cache *redis.Client) Service {

	return &service{
		config: config,
		cache:  cache,
		apiKey: config.NovaPoshtaKey,
	}
}
