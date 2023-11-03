package router

import (
	"bags2on/delivery/internal/services/shared"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Services interface {
	shared.UseCase
}

func NewRouter(sharedServices shared.UseCase) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Cookie", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Get("/popular-cities", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(sharedServices.PopularCities())
	})

	return r
}
