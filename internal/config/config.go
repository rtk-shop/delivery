package config

import (
	"rtk/delivery/internal/utils"
	"log"
	"strconv"
)

type Config struct {
	Port          string
	NovaPoshtaKey string
	NovaPoshtaURL string
	RedisDB       int
	RedisPort     string
}

func New() *Config {

	port := utils.GetEnv("PORT")
	novaPoshtaKey := utils.GetEnv("NOVA_POSHTA_KEY")
	novaPoshtaURL := utils.GetEnv("NOVA_POSHTA_API_URL")
	redisDBRaw := utils.GetEnv("REDIS_DB")
	redisPort := utils.GetEnv("REDIS_PORT")

	redisDB, err := strconv.Atoi(redisDBRaw)
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		Port:          port,
		NovaPoshtaKey: novaPoshtaKey,
		NovaPoshtaURL: novaPoshtaURL,
		RedisDB:       redisDB,
		RedisPort:     redisPort,
	}
}
