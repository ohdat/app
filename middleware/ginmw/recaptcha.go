package ginmw

import (
	"github.com/gin-gonic/gin"
	"github.com/ohdat/app/recaptcha"
	"github.com/ohdat/app/response"
)

func Recaptcha(secret string) gin.HandlerFunc {
	recaptcha.Init(secret)
	return func(c *gin.Context) {
		token := getRecaptcha(c)
		if token == "" {
			response.ErrorResponse(c, response.ErrRecaptchaNotFound)
			c.Abort()
		}
		if ok, err := recaptcha.Confirm(c.ClientIP(), token); !ok {
			response.ErrorResponse(c, err)
			c.Abort()
		}
		c.Next()
	}
}

func getRecaptcha(c *gin.Context) (token string) {
	// 客户端携带Token有三种方式 1.放在请求头  2.放在URI 3.放在请求体
	token = c.Param("recaptcha_token")
	if token != "" {
		return
	}
	token = c.Query("recaptcha_token")
	if token != "" {
		return
	}
	token = c.Request.Header.Get("recaptcha_token")
	return
}
