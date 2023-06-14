package ws

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"strconv"
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

// Publish 瓜皮兔通用消息推送.
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

// PublishMJ 用于瓜皮兔 midjourney 推送消息.
//
//	aid: 用户id.
//	nonce: 唯一字符串 雪花ID.
//	message: 消息主体.
//
// Example:
//
//	var wsPush = ws.NewGuapituPush(common.Redis).
//	wsPush.PublishMJ(1, "nonce", "message").
func (s *GuapituPush) PublishMJ(aid int, nonce string, message any) {
	var event = "mj:" + nonce
	s.Publish(aid, event, message)
}

// PublishToken 用于瓜皮兔 token 推送消息.
//
//	aid: 用户id.
//	token: 用户token.
//
// Example:
//
//	var wsPush = ws.NewGuapituPush(common.Redis)
//	wsPush.PublishToken(1, 20000)
func (s *GuapituPush) PublishToken(aid, token int) {
	s.Publish(aid, "gpt_token", token)
}

func (s *GuapituPush) PublishPayOK(aid, orderId int) {
	s.Publish(aid, "pay_success", strconv.Itoa(orderId))
}
