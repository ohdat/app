package utils

import (
	"testing"
)

func TestBalance(t *testing.T) {
	var Svc = NewEther("wss://mainnet.infura.io/ws/v3/8eea62506d27474cb952f6930fead8f2", 10)
	b, e := Svc.Balance("0xd074a09626de0aAD5391e3dC2DC5535DaD8433F0")
	t.Log(b, e)
}
