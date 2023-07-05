package ginmw

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/ohdat/app/response"
	"github.com/redis/go-redis/v9"
)

// RateByAid Limit the number of visits per second through the user's id
func RateByAid(key string, limit int, redis *redis.Client) func(c *gin.Context) {
	limiter := redis_rate.NewLimiter(redis)
	key = "_rate:" + key + ":"
	return func(c *gin.Context) {
		var aid = c.GetInt("_aid")
		var preKey = key + strconv.Itoa(aid)
		var ctx = context.Background()
		res, err := limiter.Allow(ctx, preKey, redis_rate.PerSecond(limit))
		if err != nil {
			response.ErrorResponse(c, response.ErrTooManyRequests)
			c.Abort()
			return
		}
		if res.Allowed == 0 {
			response.ErrorResponse(c, response.ErrTooManyRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}

// RateByIp Limit the number of visits per second through the user's ip
func RateByIp(key string, limit int, redis *redis.Client) func(c *gin.Context) {
	limiter := redis_rate.NewLimiter(redis)
	key = "_rate:" + key + ":"
	return func(c *gin.Context) {
		var ip = c.ClientIP()
		var preKey = key + ip
		var ctx = context.Background()
		res, err := limiter.Allow(ctx, preKey, redis_rate.PerSecond(limit))
		if err != nil {
			response.ErrorResponse(c, response.ErrTooManyRequests)
			c.Abort()
			return
		}
		if res.Allowed == 0 {
			response.ErrorResponse(c, response.ErrTooManyRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}
