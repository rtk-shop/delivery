package handlers

import (
	"encoding/json"
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

		encoder := json.NewEncoder(w)
		err = encoder.Encode(warehouses)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return

	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"message": "provider or city are wrong"}`))
}
