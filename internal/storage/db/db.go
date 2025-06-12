package db

import (
	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/akhilsomanvs/weather-api/internal/storage"
	rediscache "github.com/akhilsomanvs/weather-api/internal/storage/db/redis"
)

type Database struct {
	Storage storage.Storage
}

func InitCache(cfg *config.Config) Database {
	storage, err := rediscache.InitRedisCache(cfg)
	if err != nil {
		return Database{}
	}
	return Database{
		Storage: storage,
	}
}
