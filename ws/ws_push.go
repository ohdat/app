package ws

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

/**
* use
* ```go
* var wsPush = ws.NewWssPush(redis)
* ```
 */

// Deprecated: app v1.3.0 之后不推荐使用 预计在 v1.4.0 移除 请使用 ws.NewGuapituPush
type WssPush struct {
	Redis  *redis.Client
	SubKey string
}

/**
* use
* ```go
* var wsPush = ws.NewWssPush(redis)
* ```
 */
// Deprecated: app v1.3.0 之后不推荐使用 预计在 v1.4.0 移除 请使用 ws.NewGuapituPush
func NewWssPush(Redis *redis.Client) *WssPush {
	var subKey = viper.GetString("chatgpt.subscribe_key")
	return &WssPush{
		Redis:  Redis,
		SubKey: subKey,
	}
}

// Deprecated: app v1.3.0 之后不推荐使用 预计在 v1.4.0 移除 请使用 ws.NewGuapituPush
func (s *WssPush) Publish(message []byte) {
	ctx := context.Background()
	s.Redis.Publish(ctx, s.SubKey, message)
}

/*
  - use
    ```go
    var wsPush = ws.NewWssPush(redis)
    wsPush.PublishToken(1, 2)
    ```
*/
// Deprecated: app v1.3.0 之后不推荐使用 预计在 v1.4.0 移除 请使用 ws.NewGuapituPush
func (s *WssPush) PublishToken(aid, token int, event int, balance int) {
	accountInfo := struct {
		Aid     int `json:"aid"`
		Token   int `json:"token"`
		Event   int `json:"event"`
		Balance int `json:"balance"`
	}{
		Aid:     aid,
		Token:   token,
		Event:   event,
		Balance: balance,
	}
	bytes, _ := json.Marshal(accountInfo)
	s.Publish(bytes)
}
