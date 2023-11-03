package handlers

import "net/http"

func (h *Handlers) PopularCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(h.sharedUC.PopularCities())
}
