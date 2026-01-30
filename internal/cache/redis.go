package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(addr, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Println("Failed to pin Redis")
		return nil, err
	}
	return rdb, nil
}
