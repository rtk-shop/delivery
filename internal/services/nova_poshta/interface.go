package novaposhta

import (
	"bags2on/delivery/internal/config"

	"github.com/redis/go-redis/v9"
)

type UseCase interface {
	Warehouses(cityID string) ([]byte, error)
}

type service struct {
	config *config.Config
	cache  *redis.Client
}

func NewNovaPoshtaService(config *config.Config, cache *redis.Client) UseCase {

	return &service{
		config: config,
		cache:  cache,
	}
}
