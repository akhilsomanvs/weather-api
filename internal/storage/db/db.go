package db

import (
	"encoding/json"
	"fmt"

	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/models"
	"github.com/akhilsomanvs/weather-api/internal/storage"
	rediscache "github.com/akhilsomanvs/weather-api/internal/storage/db/redis"
)

type Database struct {
	Storage storage.Storage
}

func InitCache(cfg *config.Config) (*Database, error) {
	storage, err := rediscache.InitRedisCache(cfg)
	if err != nil {
		return &Database{}, fmt.Errorf("Could not create instance of Cache server: %s", err.Error())
	}
	return &Database{
		Storage: storage,
	}, nil
}

func (db Database) GetWeatherData(key string) (*models.WeatherApiResponseModel, error) {
	val, err := db.Storage.GetData(key)
	if err != nil {
		return &models.WeatherApiResponseModel{}, fmt.Errorf("Could not fetch data from cache: %s", err.Error())
	}

	var response models.WeatherApiResponseModel
	err = json.Unmarshal([]byte(val.(string)), &response)
	if err != nil {
		return &models.WeatherApiResponseModel{}, fmt.Errorf("Could not fetch data from cache: %s", err.Error())
	}
	return &response, nil
}

func (db Database) SetWeatherData(key string, model models.WeatherApiResponseModel) error {
	data, err := json.Marshal(model)
	if err != nil {
		return fmt.Errorf("Could not store data to cache: %s", err.Error())
	}
	err = db.Storage.SetData(key, data)
	if err != nil {
		return fmt.Errorf("Could not store data to cache: %s", err.Error())
	}
	return nil
}
