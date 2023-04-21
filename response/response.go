package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message,omitempty"`
	Timestamp int64       `json:"timestamp,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

const (
	// CodeSuccess
	CodeSuccess = 200
	// CodeCommon
	CodeCommon = 1001
)

// ErrorResponse
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

// SuccessResponse
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code:      CodeSuccess,
		Message:   "success",
		Timestamp: time.Now().Unix(),
		Data:      data,
	})
}

func SuccessJson(data interface{}) ([]byte, error) {
	return json.Marshal(&Response{
		Code:      CodeSuccess,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
	})
}
