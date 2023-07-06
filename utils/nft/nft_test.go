package nft

import (
	"log"
	"testing"
)

func TestTotal(t *testing.T) {
	total, _ := NewNft("LZbjN9bMxemRAARcIVgnG3p3I7lVCt7E").Total("0xfc7c3b0fd8e183b499149e9efc083af7ba43d720")
	//t.Log("Transfer721And20TopicHash:", Transfer721And20TopicHash)
	log.Println("total:", total)
}

func TestFormTransfer(t *testing.T) {
	NewNft("LZbjN9bMxemRAARcIVgnG3p3I7lVCt7E").FormTransfer("0xfc7c3b0fd8e183b499149e9efc083af7ba43d720", "")
}

func TestTransferTotal(t *testing.T) {
	total := NewNft("LZbjN9bMxemRAARcIVgnG3p3I7lVCt7E").TransferTotal("0xfc7c3b0fd8e183b499149e9efc083af7ba43d720")
	log.Println("total:", total)
}
