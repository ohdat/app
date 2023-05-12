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
func NewWssPush(Redis *redis.Client) *WssPush {
	var subKey = viper.GetString("chatgpt.subscribe_key")
	return &WssPush{
		Redis:  Redis,
		SubKey: subKey,
	}
}

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
func (s *WssPush) PublishToken(aid, token int) {
	accountInfo := &AccountInfo{
		Aid:   aid,
		Token: token,
	}
	bytes, _ := json.Marshal(accountInfo)
	s.Publish(bytes)
}
