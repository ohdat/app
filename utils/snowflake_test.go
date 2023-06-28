package utils

import (
	"log"
	"testing"
)

func TestNextSnowflakeID(t *testing.T) {
	id := NextSnowflakeID()
	t.Log(id)
}

func TestDCNonce(t *testing.T) {
	id := DCNonce()
	log.Println("id", id)
	t.Log(id)
}
