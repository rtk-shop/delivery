package novaposhta

import (
	"rtk/delivery/internal/config"
	"rtk/delivery/internal/entity"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	Warehouses(cityID string) ([]byte, error)
	Settlements(cityName string) ([]entity.NovaPoshtaSettlement, error)
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
