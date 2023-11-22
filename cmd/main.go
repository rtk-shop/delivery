package main

import (
	"bags2on/delivery/internal/app"
	"bags2on/delivery/internal/config"
	"bags2on/delivery/pkg/cache"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env", ".env.secret"); err != nil {
		log.Fatalf(".env load error: %v", err)
	}
	log.Printf("loaded .env files\n")
}

func main() {
	config := config.New()
	cache := cache.NewRedisClient(config)

	app := app.New(config, cache)

	app.Run()
}
