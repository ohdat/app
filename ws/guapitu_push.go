package ws

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// GuapituPushMessage 瓜皮兔推送消息结构体.
type GuapituPushMessage struct {
	Aid     int    `json:"aid"`
	Event   string `json:"event"`
	Message any    `json:"message"`
}

// GuapituPush 用于瓜皮兔推送消息.
//
//	Example:
//
//	var wsPush = ws.NewGuapituPush(common.Redis)
type GuapituPush struct {
	Redis  *redis.Client
	SubKey string
}

// NewGuapituPush 用于瓜皮兔推送消息.
//
//	Example:
//
// var wsPush = ws.NewGuapituPush(redis)
func NewGuapituPush(Redis *redis.Client) *GuapituPush {
	var subKey = viper.GetString("chatgpt.ws_sub_key")
	return &GuapituPush{
		Redis:  Redis,
		SubKey: subKey,
	}
}

func (s *GuapituPush) publish(message []byte) {
	ctx := context.Background()
	s.Redis.Publish(ctx, s.SubKey, message)
}

// Publish 用于瓜皮兔推送消息.
//
//	aid: 用户id.
//	event: 事件.
//	message: 消息主体.
//
// Example:
//
//	var wsPush = ws.NewGuapituPush(redis).
//	wsPush.Publish(1, "event", "message").
func (s *GuapituPush) Publish(aid int, event string, message any) {
	eventMsg := GuapituPushMessage{
		Aid:     aid,
		Event:   event,
		Message: message,
	}
	bytes, _ := json.Marshal(eventMsg)
	s.publish(bytes)
}

func (s *GuapituPush) PublishMJ(aid int, message any) {
	var event = "mj"
	s.Publish(aid, event, message)
}
