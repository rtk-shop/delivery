package novaposhta

import "bags2on/delivery/internal/config"

type UseCase interface {
	Warehouses(cityID string) ([]byte, error)
}

type service struct {
	config *config.Config
}

func NewNovaPoshtaService(config *config.Config) UseCase {
	return &service{
		config: config,
	}
}
