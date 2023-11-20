package redis

import (
	"bags2on/delivery/internal/config"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewClient(config *config.Config) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:" + config.RedisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println("redis connection not established")
		panic(err)
	}

	return client
}
