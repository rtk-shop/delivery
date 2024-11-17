package novaposhta

import (
	"rtk/delivery/internal/config"
	"rtk/delivery/internal/entity"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	Warehouses(cityID string, warehouseType int) ([]entity.NovaPoshtaWarehouse, error)
	Settlements(cityName string) ([]entity.NovaPoshtaSettlement, error)
}

type service struct {
	config *config.Config
	cache  *redis.Client
	apiKey string
}

var warehouseTypesMap = map[int]string{
	1: "841339c7-591a-42e2-8233-7a0a00f0ed6f", // Почтовое отделение
	2: "9a68df70-0267-42a8-bb5c-37f427e36ee4", // Грузовое отделение
	3: "f9316480-5f2d-425d-bc2c-ac7cd29decf0", // Почтомат

}

func New(config *config.Config, cache *redis.Client) Service {

	return &service{
		config: config,
		cache:  cache,
		apiKey: config.NovaPoshtaKey,
	}
}
