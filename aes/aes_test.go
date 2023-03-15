package aes

import (
	"testing"
)

func TestCrypto(t *testing.T) {

	var aes = New("TbeRlq2aXl2nTmHZTbeRlq2aXl2nTmHZ")

	ciphertext, err := aes.Encode("mysqlpassword")

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ciphertext: %s\n", ciphertext)

	decryptMessage, err := aes.Decode(ciphertext)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("decryptMessage: %s\n", decryptMessage)
}
