package app

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"rtk/delivery/internal/config"
	"rtk/delivery/internal/delivery/router"
	"rtk/delivery/internal/delivery/router/handlers"

	novaposhta "rtk/delivery/internal/services/nova-poshta"
	"rtk/delivery/internal/services/shared"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type app struct {
	config *config.Config
	router chi.Router
}

func New(config *config.Config, logger *slog.Logger, cache *redis.Client) *app {

	sharedService := shared.New(config, cache)

	nvpService := novaposhta.New(config, logger, cache)

	handlers := handlers.NewHandlers(sharedService, nvpService)

	router := router.NewRouter(handlers)

	return &app{
		config: config,
		router: router,
	}
}

func (a *app) Run() {
	a.serve()
}

func (a *app) serve() {
	server := &http.Server{
		Addr:              ":" + a.config.Port,
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       1 * time.Second,
		Handler:           http.TimeoutHandler(a.router, 15*time.Second, "request timeout expired"),
	}

	log.Printf("try to connect to http://localhost:%s", a.config.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("server %q - shutting down", <-done)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("❌ server shutdown failed:%+v", err)
	}

	log.Print("✅ server shutdown gracefully")
}
