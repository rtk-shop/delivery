package novaposhta

import (
	"log/slog"
	"rtk/delivery/internal/config"
	"rtk/delivery/internal/entity"

	"github.com/redis/go-redis/v9"
)

/*
	Почему []byte а не entity.NovaPoshtaWarehouse?

	Бенчмарк запроса всех почтоматов в Киеве может достигать ~2 мб
	и marshal/unmarshal потребляет слишком много,
	не говоря о дополнительной сериализации в транспортном слое

	BenchmarkWarehousesJSON-8   	     223	    5015052 ns/op	  3278636 B/op	   14419 allocs/op
	BenchmarkWarehousesByte-8   	   11002	     107362 ns/op	   999700 B/op	       9 allocs/op
*/

type Service interface {
	Warehouses(cityID string, warehouseType int) ([]byte, error)
	Settlements(cityName string) ([]entity.NovaPoshtaSettlement, error)
}

type service struct {
	config *config.Config
	logger *slog.Logger
	cache  *redis.Client
	apiKey string
}

var warehouseTypesMap = map[int]string{
	1: "841339c7-591a-42e2-8233-7a0a00f0ed6f", // Почтовое отделение
	2: "9a68df70-0267-42a8-bb5c-37f427e36ee4", // Грузовое отделение
	3: "f9316480-5f2d-425d-bc2c-ac7cd29decf0", // Почтомат

}

func New(config *config.Config, logger *slog.Logger, cache *redis.Client) Service {

	return &service{
		config: config,
		logger: logger,
		cache:  cache,
		apiKey: config.NovaPoshtaKey,
	}
}
