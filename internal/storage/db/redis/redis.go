package rediscache

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/akhilsomanvs/weather-api/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	Client *redis.Client
	Config *config.Config
}

func InitRedisCache(cfg *config.Config) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.GetCacheHostAddress(),
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Could not connect to Cache server : %s", err.Error())
		return nil, err
	}
	return &RedisCache{
		Client: client,
		Config: cfg,
	}, nil
}

func (cache *RedisCache) SetData(key string, data any) error {
	err := cache.Client.Set(context.Background(), key, data, time.Duration(rand.Int31n(cache.Config.CacheDuration))*time.Hour).Err()
	if err != nil {
		log.Printf("Could not store data in cache: %s", err.Error())
		return err
	}
	return nil
}

func (cache *RedisCache) GetData(key string) (any, error) {
	val, err := cache.Client.Get(context.Background(), key).Result()
	if err != nil {
		log.Printf("Could not fetch data from cache: %s", err.Error())
		return nil, err
	}

	return val, nil
}
