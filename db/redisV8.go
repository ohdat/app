package db

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
	"sync"
)

func initRedisV8(db int) *redis.Client {
	opts := redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           db,
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: viper.GetInt("redis.min_idle_conns"),
	}

	client := redis.NewClient(&opts)

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalln("RedisV8 Init Error: ", err)
	}
	return client
}

var redisV8 *redis.Client
var redisOnce sync.Once

func GetRedis() *redis.Client {
	redisOnce.Do(func() {
		redisV8 = initRedisV8(viper.GetInt("redis.db"))
	})
	return redisV8
}

var durableOnce sync.Once
var durableRedis *redis.Client

func GetDurableRedis() *redis.Client {
	durableOnce.Do(func() {
		durableRedis = initRedisV8(viper.GetInt("redis.durable_db"))
	})
	return durableRedis
}
