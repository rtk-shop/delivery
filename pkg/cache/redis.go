package cache

import (
	"bags2on/delivery/internal/config"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config *config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + config.RedisPort,
		Password: "",
		DB:       config.RedisDB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("redis connection not established")
		panic(err)
	}

	return client
}
