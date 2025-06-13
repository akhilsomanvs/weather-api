package storage

import "github.com/akhilsomanvs/weather-api/internal/models"

type Storage interface {
	SetData(key string, data []byte) error
	GetData(key string) (any, error)
}

type WeatherDBController interface {
	GetWeatherData(key string) (*models.WeatherApiResponseModel, error)
	SetWeatherData(key string, model models.WeatherApiResponseModel) error
}
