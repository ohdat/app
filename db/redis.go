package db

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func initRedis(db int) *redis.Client {
	opts := redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           db,
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: viper.GetInt("redis.min_idle_conns"),
	}

	client := redis.NewClient(&opts)

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("redis Init Error: ", err)
	}
	return client
}

var appRedis *redis.Client
var redisOnce sync.Once

func GetRedis() *redis.Client {
	redisOnce.Do(func() {
		appRedis = initRedis(viper.GetInt("redis.db"))
	})
	return appRedis
}
