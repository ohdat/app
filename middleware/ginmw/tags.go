package ginmw

import (
	"github.com/gin-gonic/gin"
	"github.com/ohdat/app/middleware/tags/gintags"
)

func Tags() gin.HandlerFunc {
	return func(c *gin.Context) {
		gintags.Set(c, "test.tags", "c.Request.URL.Path")
		c.Next()
	}
}
