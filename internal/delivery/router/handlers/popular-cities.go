package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func (h *Handlers) PopularCities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cities, err := h.sharedService.PopularCities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err)))
		return
	}

	eTag := h.sharedService.GetPopularCitiesHash()

	// Remove quotes and W/ prefix for If-None-Match header comparison
	ifNoneMatch := strings.TrimPrefix(strings.Trim(r.Header.Get("If-None-Match"), "\""), "W/")

	// Generate a hash of the content without the W/ prefix for comparison
	contentHash := strings.TrimPrefix(eTag, "W/")

	// Check if the ETag matches; if so, return 304 Not Modified
	if ifNoneMatch == strings.Trim(contentHash, "\"") {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	w.Header().Set("ETag", eTag) // Send weak ETag
	// w.Header().Set("Cache-Control", "max-age=604800") // one week
	w.Write(cities)
}
