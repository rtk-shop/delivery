package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handlers) PopularCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cities, err := h.sharedUC.PopularCities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err)))
		return
	}

	w.Write(cities)
}
