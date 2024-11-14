package main

import (
	"flag"
	"log"
	"rtk/delivery/internal/app"
	"rtk/delivery/internal/config"
	"rtk/delivery/pkg/cache"

	"github.com/joho/godotenv"
)

func init() {
	env := flag.String("env", "dev", "specify .env filename for flag")
	flag.Parse()

	if *env == "dev" {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Fatal("No .env.local file found, create it!")
		}
		log.Println("loaded .env.local")
	}

	if err := godotenv.Load(".env." + *env); err != nil {
		log.Fatalf("No .env.%s file found, load default", *env)
	}

	log.Printf("loaded \".env.%s\"\n", *env)

}

func main() {
	config := config.New()

	cache := cache.NewRedisClient(config)

	app := app.New(config, cache)

	app.Run()
}
