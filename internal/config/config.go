package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ENV       string    `yaml:"env"`
	ApiConfig ApiConfig `yaml:"api_config"`
}

type ApiConfig struct {
	ApiKey  string `yaml:"API_KEY"`
	BaseUrl string `yaml:"API_BASE_URL"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "Path to configuration file")
		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatalln("Config path is not set")
		}
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file : %s", err.Error())
	}

	return &cfg

}
