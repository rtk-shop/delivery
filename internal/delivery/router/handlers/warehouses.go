package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handlers) Warehouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	provider := strings.TrimSpace(r.URL.Query().Get(ProviderKey))
	cityId := strings.TrimSpace(r.URL.Query().Get(CityKey))
	warehouseType := strings.TrimSpace(r.URL.Query().Get(WarehouseTypeKey))

	if provider == NovaProvider {

		warehouseTypeId, err := strconv.Atoi(warehouseType)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"message": "%s should be a number"}`, WarehouseTypeKey)
			return
		}

		warehouses, err := h.nvpService.Warehouses(cityId, warehouseTypeId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"message": "%s"}`, err)
			return
		}

		_, err = w.Write(warehouses)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, `{"message": "failed to send response"}`)
			return
		}

		return

	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"message": "%s is wrong"}`, ProviderKey)
}
