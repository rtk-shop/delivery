package handlers

import (
	// novaposhta "rtk/delivery/internal/services/nova_poshta"
	novaposhta "rtk/delivery/internal/services/nova-poshta"
	"rtk/delivery/internal/services/shared"
)

type Handlers struct {
	sharedService shared.Service
	nvpService    novaposhta.Service
}

func NewHandlers(shared shared.Service, novaposhta novaposhta.Service) *Handlers {
	return &Handlers{
		sharedService: shared,
		nvpService:    novaposhta,
	}
}
