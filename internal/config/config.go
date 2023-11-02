package config

import "bags2on/delivery/internal/utils"

type Config struct {
	Port          string
	NovaPoshtaKey string
}

func New() *Config {

	port := utils.GetEnv("PORT")
	novaPoshtaKey := utils.GetEnv("NOVA_POSHTA_KEY")

	return &Config{
		Port:          port,
		NovaPoshtaKey: novaPoshtaKey,
	}
}
