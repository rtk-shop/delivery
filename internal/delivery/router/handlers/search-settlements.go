package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (h *Handlers) SearchSettlements(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	provider := strings.TrimSpace(r.URL.Query().Get(ProviderKey))
	cityName := strings.TrimSpace(r.URL.Query().Get(CityNameKey))

	if provider == "" || cityName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"message": "%s or %s are wrong"}`, ProviderKey, CityNameKey)
		return
	}

	if provider == NovaProvider {

		settlements, err := h.nvpService.Settlements(cityName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err)))
			return
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(settlements)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, `{"message": "%s or %s are wrong"}`, ProviderKey, CityNameKey)

}
