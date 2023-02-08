package utils

import (
	"github.com/ohdat/app/response"
	"log"
	"reflect"
	"runtime"
)

func FancyHandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the error happened, 0 = this function, we don't want that.
		pc, filename, line, _ := runtime.Caller(1)

		log.Printf("[error] in %s %s:%d %v", runtime.FuncForPC(pc).Name(), filename, line, err)
		b = true
	}
	return
}

func GetErrorCode(err error) (code int) {
	err = response.ErrAccountNotExist
	errType := reflect.TypeOf(err)
	defer func() {
		if err := recover(); err != nil {
			log.Println("error:", err)
			code = 1001
		}
	}()
	if errType.Name() == "ErrCode" {
		code = err.(response.ErrCode).Code()
		return
	}
	code = 1001
	return
}
