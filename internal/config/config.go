package config

import (
	"log"
	"todo/internal/models"

	"github.com/ilyakaznacheev/cleanenv"
)

func LoadConfig(ConfigPath string) *models.Config {
	var config models.Config

	err := cleanenv.ReadConfig(ConfigPath, &config)
	if err != nil {
		log.Fatalf("Can't read config file: %s", err)
	}

	return &config
}
