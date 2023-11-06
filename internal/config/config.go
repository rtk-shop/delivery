package config

import "bags2on/delivery/internal/utils"

type Config struct {
	Port          string
	NovaPoshtaKey string
	NovaPoshtaURL string
}

func New() *Config {

	port := utils.GetEnv("PORT")
	novaPoshtaKey := utils.GetEnv("NOVA_POSHTA_KEY")
	novaPoshtaURL := utils.GetEnv("NOVA_POSHTA_API_URL")

	return &Config{
		Port:          port,
		NovaPoshtaKey: novaPoshtaKey,
		NovaPoshtaURL: novaPoshtaURL,
	}
}
