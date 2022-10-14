package gintags

import (
	"github.com/gin-gonic/gin"
	"sync"
)

const (
	ginCtxTagsKey = "gin-ctx-tags"
)

var mu sync.RWMutex

func Set(ctx *gin.Context, name string, value interface{}) {
	mu.Lock()
	tags := Values(ctx)
	tags[name] = value
	ctx.Set(ginCtxTagsKey, tags)
	mu.Unlock()
}

func Values(ctx *gin.Context) map[string]interface{} {
	tags, ok := ctx.Get(ginCtxTagsKey)
	if !ok {
		return make(map[string]interface{})
	}
	return tags.(map[string]interface{})
}
