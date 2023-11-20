package main

import (
	"bags2on/delivery/internal/app"
	"bags2on/delivery/internal/config"
	"bags2on/delivery/pkg/redis"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("no .env files found")
	}
	log.Printf("loaded .env\n")
}

func main() {
	config := config.New()
	cache := redis.NewClient(config)

	app := app.New(config, cache)

	app.Run()
}
