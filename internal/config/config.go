package config

type Config struct {
	ENV       string    `yaml:"env"`
	ApiConfig ApiConfig `yaml:"api_config"`
}

type ApiConfig struct {
	ApiKey  string `yaml:"API_KEY"`
	BaseUrl string `yaml:"API_BASE_URL"`
}
