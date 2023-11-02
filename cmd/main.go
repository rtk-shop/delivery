package main

import (
	"bags2on/delivery/internal/app"
	"bags2on/delivery/internal/config"
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

	app := app.New(config)

	app.Run()
}
