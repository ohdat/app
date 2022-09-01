package aes

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	var sss = New("oOC)dO9IFoExp&nzo$gr*$X4xl*qd(($")

	ccc, err := sss.Encode("1111.jpg")
	if err != nil {
		t.Error(err)
	}
	bbb, err := sss.Decode(ccc)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("ccc", ccc)
	fmt.Println("bbb", bbb)
	t.Log("ccc", ccc)
	t.Log("bbb", bbb)
}
