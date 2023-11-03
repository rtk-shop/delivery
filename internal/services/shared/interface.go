package shared

import "bags2on/delivery/internal/config"

type UseCase interface {
	PopularCities() []byte
}

type service struct {
	config *config.Config
}

func NewSharedService(config *config.Config) UseCase {
	return &service{
		config: config,
	}
}
