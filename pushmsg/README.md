# github.com/ohdat/app/pushmsg

GuapituPush 用于瓜皮兔推送消息.

```go
type GuapituPush struct {
    Redis  *redis.Client
    SubKey string
}
# 消息体
Event struct {
    Event   string      `json:"event"`
    Message any         `json:"message"`
}

```

## front end 
```js
{
    event: 'event',
    message: 'message'
}
```

### 常用 event
1. `mj:${nonce}` midjourney 消息
2. `gpt_token`  瓜皮兔token 变化
3. `gpt:pay:success:{orderId}` 支付成功
4. `gpt:vip:level` 等级变更



## How to use?

####  通用推送消息
```go
// import "github.com/ohdat/app/ws"
    var wsPush = ws.NewGuapituPush(common.Redis)
    wsPush.Publish(1, "event", "message").
```

####  推送 midjourney 消息
```go
// import "github.com/ohdat/app/ws"
    var wsPush = ws.NewGuapituPush(common.Redis)
    wsPush.PublishMJ(1, "once", "message").
```

## 开发规范

新增需求的时候尽量不要修改之前函数的 参数和返回值。尽可能采用新增的方式。保证之前的函数不受影响。  
如果之前的函数需要修改，在注释中添加 Deprecated 标记。并且在函数体中调用新的函数。
