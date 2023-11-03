package app

import (
	"bags2on/delivery/internal/config"
	"bags2on/delivery/internal/services/shared"
)

func SharedService(config *config.Config) shared.UseCase {

	return shared.NewSharedService(config)
}

type ServicesRoot struct {
	Shared shared.UseCase
}
