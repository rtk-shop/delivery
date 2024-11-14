package handlers

import (
	novaposhta "rtk/delivery/internal/services/nova_poshta"
	"rtk/delivery/internal/services/shared"
)

type Handlers struct {
	sharedUC     shared.UseCase
	novaposhtaUC novaposhta.UseCase
}

func NewHandlers(shared shared.UseCase, novaposhtaUC novaposhta.UseCase) *Handlers {
	return &Handlers{
		sharedUC:     shared,
		novaposhtaUC: novaposhtaUC,
	}
}
