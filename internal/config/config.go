package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ENV        string `yaml:"env"`
	ApiConfig  `yaml:"api_config"`
	HttpServer `yaml:"http_server"`
	Cache      `yaml:"cache"`
}

type ApiConfig struct {
	ApiKey  string `yaml:"API_KEY"`
	BaseUrl string `yaml:"API_BASE_URL"`
}

type HttpServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Cache struct {
	CacheHost     string `yaml:"host"`
	CachePort     string `yaml:"port"`
	CacheDuration int32  `yaml:"duration_in_hours"`
}

func (cfg *Config) GetHostAddress() string {
	return (cfg.Host + ":" + cfg.Port)
}

func (cfg *Config) GetCacheHostAddress() string {
	return (cfg.CacheHost + ":" + cfg.CachePort)
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
