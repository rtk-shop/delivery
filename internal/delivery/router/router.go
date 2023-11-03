package router

import (
	"bags2on/delivery/internal/delivery/router/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(handlers *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Cookie", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	}))

	r.Get("/popular-cities", handlers.PopularCities)
	r.Get("/warehouses", handlers.Warehouses)

	return r
}
