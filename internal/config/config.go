package config

import "bags2on/delivery/internal/utils"

type Config struct {
	Port          string
	NovaPoshtaKey string
	NovaPoshtaURL string
	RedisPort     string
}

func New() *Config {

	port := utils.GetEnv("PORT")
	novaPoshtaKey := utils.GetEnv("NOVA_POSHTA_KEY")
	novaPoshtaURL := utils.GetEnv("NOVA_POSHTA_API_URL")
	redisPort := utils.GetEnv("REDIS_PORT")

	return &Config{
		Port:          port,
		NovaPoshtaKey: novaPoshtaKey,
		NovaPoshtaURL: novaPoshtaURL,
		RedisPort:     redisPort,
	}
}
