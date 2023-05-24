# ws

GuapituPush 用于瓜皮兔推送消息.

```go
type GuapituPush struct {
        Redis  *redis.Client
        SubKey string
}
```

## How to use?

```go
// import "github.com/ohdat/app/ws"
    var wsPush = ws.NewGuapituPush(common.Redis)
    wsPush.Publish(1, "event", "message").
```

## 开发规范

新增需求的时候尽量不要修改之前函数的 参数和返回值。尽可能采用新增的方式。保证之前的函数不受影响。  
如果之前的函数需要修改，在注释中添加 Deprecated 标记。并且在函数体中调用新的函数。
