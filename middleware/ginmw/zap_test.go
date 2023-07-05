package ginmw

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ohdat/app/logger"
	"github.com/ohdat/app/tags/gintags"
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
	})
	go server.Run(":8080")
	time.Sleep(time.Second * 1)
	const url = "http://127.0.0.1:8080/ping?de=1"
	body := []byte(`{"name":"test"}`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("error occurred while sending request to %s: %v", url, err)
	}
	defer resp.Body.Close()

	log.Printf("response status: %s", resp.Status)

	// read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response body: %v", err)
	}
	log.Printf("response body: %s", respBody)
}
