package pkg

import (
	"github.com/redis/go-redis/v9"
	"sync"
)

type RedisClient struct {
	Client *redis.Client
}

var once sync.Once

var Redis = NewRedisClient()

func ConnectRedis() {
	Redis = NewRedisClient()
}

func CreateRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func NewRedisClient() *RedisClient {
	return &RedisClient{
		Client: CreateRedisClient(),
	}
}

func (rc *RedisClient) Test() {}
