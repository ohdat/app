package utils

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	appredis "github.com/redis/go-redis/v9"
)

const keyPrefix = "lock:"

// GetLock 获取锁
// how to use
// lock := GetLock(db.RedisV8, fmt.Sprintf("ticket_balance_reduce:%d", id))
// err = lock.Lock()
//
//	if err != nil {
//		fmt.Println("xxx err",err)
//	}
//
// defer  lock.Unlock()
func GetLock(redisv9 *appredis.Client, key string) (mutex *redsync.Mutex) {
	pool := goredis.NewPool(redisv9)
	rs := redsync.New(pool)
	return rs.NewMutex(keyPrefix + key)
}
