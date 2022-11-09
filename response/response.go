package response

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

// Response 返回结构
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

const (
	// CodeSuccess 正确
	CodeSuccess = 200
	// CodeCommon 通用错误
	CodeCommon = 1001
)

// ErrorResponse 返回错误
func ErrorResponse(c *gin.Context, infos ...interface{}) {
	var err = infos[0].(error)
	response := &Response{
		Code:      CodeCommon,
		Message:   err.Error(),
		Timestamp: time.Now().Unix(),
	}

	if len(infos) > 1 {
		response.Data = infos[1]
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
		c.JSON(http.StatusOK, response)
	}()
	errType := reflect.TypeOf(err)
	if errType.Name() == "ErrCode" {
		//typ catch
		response.Code = err.(ErrCode).Code()
	}
}

// SuccessResponse 返回带有数据的正确响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:      CodeSuccess,
		Message:   "OK",
		Timestamp: time.Now().Unix(),
		Data:      data,
	})
}
