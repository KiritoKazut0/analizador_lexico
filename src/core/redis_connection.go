package core

import (
	"context"
	"fmt"
	"os"
	"github.com/redis/go-redis/v9"
)

func ConnectRedisClient()(*redis.Client, error) {

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	dns := host + ":" + port
	password := os.Getenv("REDIS_PASSWORD")


	rdb := redis.NewClient(&redis.Options{
		Addr:     dns,
		Password: password,
		DB:       0,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}

	fmt.Println("Connected to Redis successfully")
	return rdb, nil

}


