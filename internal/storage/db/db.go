package db

import (
	"fmt"

	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/storage"
	rediscache "github.com/akhilsomanvs/weather-api/internal/storage/db/redis"
)

type Database struct {
	Storage storage.Storage
}

func InitCache(cfg *config.Config) (Database, error) {
	storage, err := rediscache.InitRedisCache(cfg)
	if err != nil {
		return Database{}, fmt.Errorf("Could not create instance of Cache server: %s", err.Error())
	}
	return Database{
		Storage: storage,
	}, nil
}
