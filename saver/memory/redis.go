package memory

import (
	"context"
	"log"
	"saver/config"

	"github.com/redis/go-redis/v9"
)

func New(config config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.Host + ":" + config.Port,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	return client
}
