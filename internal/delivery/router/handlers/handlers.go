package handlers

import (
	novaposhta "bags2on/delivery/internal/services/nova_poshta"
	"bags2on/delivery/internal/services/shared"
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
