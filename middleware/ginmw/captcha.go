package ginmw

import (
	"github.com/gin-gonic/gin"
	"github.com/ohdat/app/response"
	"github.com/ohdat/grpcclient/captcha"
)

func Captcha(grpcAddr string) gin.HandlerFunc {
	var svc, _ = captcha.NewCaptchaClient(grpcAddr)
	return func(c *gin.Context) {
		token := getCaptchaToken(c)
		if token == "" {
			response.ErrorResponse(c, response.ErrCaptchaTokenNotFound)
			c.Abort()
		}
		if err := svc.CheckToken(token, c.ClientIP()); err != nil {
			response.ErrorResponse(c, response.ErrCaptchaFailed)
			c.Abort()
		}
		c.Next()
	}
}

func getCaptchaToken(c *gin.Context) (token string) {
	// There are three ways to get a token.
	// 1. Put it in the request header.
	// 2. Put it in the URI.
	// 3. Put it in the request body.
	token = c.PostForm("_captcha_token")
	if token != "" {
		return
	}
	token = c.Query("_captcha_token")
	if token != "" {
		return
	}
	token = c.Request.Header.Get("_captcha_token")
	return
}
