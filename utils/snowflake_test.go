package utils

import (
	"log"
	"testing"
<<<<<<< HEAD
=======
	"time"

	"github.com/sony/sonyflake"
>>>>>>> d0e9b08 (upgrade redis v9)
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
