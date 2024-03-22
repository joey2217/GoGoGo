package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

const EXPIRATION = 24 * time.Hour

var redisClient *redis.Client
var redisCtx = context.Background()

func NewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisCtx = context.Background()
	redisClient = client
}

func Get(key string) (string, error) {
	return redisClient.Get(redisCtx, key).Result()
}

func Set(key string, value string) error {
	return redisClient.Set(redisCtx, key, value, EXPIRATION).Err()
}

func GetModel(key string, v any) error {
	if data, err := redisClient.Get(redisCtx, key).Result(); err == nil {
		e := json.Unmarshal([]byte(data), v)
		if e != nil {
			return e
		} else {
			return nil
		}
	} else {
		return err
	}
}

func SetModel(key string, v any) error {
	if json, err := json.Marshal(v); err == nil {
		return redisClient.Set(redisCtx, key, json, EXPIRATION).Err()
	} else {
		return err
	}
}
