package shared

import (
	"encoding/json"
	"log"
	"os"
	"rtk/delivery/internal/config"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	PopularCities() ([]byte, error)
}

type service struct {
	config         *config.Config
	cache          *redis.Client
	populdarCities []byte
}

func New(config *config.Config, cache *redis.Client) Service {

	data, err := os.ReadFile("json/popular_cities.json")
	if err != nil {
		log.Fatalln("failed to parse popular_cities.json", "error", err)
	}

	ok := json.Valid(data)
	if !ok {
		log.Fatalln("popular_cities.json not valid")
	}

	return &service{
		config:         config,
		cache:          cache,
		populdarCities: data,
	}
}
