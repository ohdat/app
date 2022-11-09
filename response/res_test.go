package response

import (
	"reflect"
	"testing"
)

func TestResponse(t *testing.T) {
	var l1 error
	l1 = ErrNeedPay
	st := reflect.TypeOf(l1)
	t.Log(st.Name())
}
