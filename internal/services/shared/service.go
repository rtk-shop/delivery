package shared

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"rtk/delivery/internal/config"
	"rtk/delivery/internal/utils"

	"github.com/redis/go-redis/v9"
)

type Service interface {
	PopularCities() ([]byte, error)
	GetPopularCitiesHash() string
}

type service struct {
	config             *config.Config
	cache              *redis.Client
	populdarCities     []byte
	populdarCitiesHash string
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

	hash, err := utils.GetMD5hash(bytes.NewReader(data))
	if err != nil {
		log.Fatalln("failed to generate MD5 for popular_cities.json", "error", err)
	}

	return &service{
		config:             config,
		cache:              cache,
		populdarCities:     data,
		populdarCitiesHash: hash,
	}
}
