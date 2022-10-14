package ginm

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ohdat/app/logger"
	"github.com/ohdat/app/tags/gintags"
	"net/http"
	"testing"
	"time"
)

func TestZap(t *testing.T) {
	var server = gin.Default()
	var zapLogger = logger.NewZapLogger("zap-test")
	server.Use(Tags(), Ginzap(zapLogger, &Config{
		RequestBody:  true,
		ResponseBody: true,
		//SkipPaths: []string{"/ping"},
	}))

	//router := server.Group("/", TraceMiddleware(zapLogger))

	server.POST("/ping", func(c *gin.Context) {
		gintags.Set(c, "auth.uid", 12)
		zapLogger.Info("test2")
		c.JSON(http.StatusOK, gin.H{
			"s": "pong",
		})
		return
	})
	go server.Run(":8080")
	time.Sleep(time.Second * 1)
	var body = `{"name":"test"}`

	response, err := http.Post("http://127.0.0.1:8080/ping?de=1", "application/json", bytes.NewBufferString(body))

	t.Log(response)
	t.Log(err)

}
