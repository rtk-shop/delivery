package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	CityKey      = "city_id"
	ProviderKey  = "provider"
	NovaProvider = "nova_poshta"
	UkrProvider  = "ukr_poshta"
)

func (h *Handlers) Warehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cityId := strings.TrimSpace(r.URL.Query().Get(CityKey))
	provider := strings.TrimSpace(r.URL.Query().Get(ProviderKey))
	// fmt.Println("query params:", cityId, provider)

	if provider == NovaProvider {
		data, err := h.nvpService.Warehouses(cityId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err)))
			return
		}
		w.Write(data)
		return

	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"message": "provider or city are wrong"}`))
}
