package main

import (
	"rtk/delivery/internal/app"
	"rtk/delivery/internal/config"
	"rtk/delivery/pkg/cache"
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
