package wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestPaymentPay(t *testing.T) {
	var s = NewSignature("0xdc4267dae71b06376c11c14063203d769aa56a4d73513e33d17ebf1c27d818ec")

	var (
		operatorAddress string   = "0xba7425B0E095da6D84B3425d295af6dfFDDdD017"
		amount          *big.Int = big.NewInt(1000000000000000000)
		tokenAddress    string   = "0xe4086f2144a2d7e8568860f142d316c8d861c2a2"
		toAddress       string   = "0xe8bB7a6BD5b7b1c4303D5b80c27c4f301768Ee03"
		nonce           string   = "iq20230213075623447524491464023404"
		hashStr         string
		err             error
		signature       string
	)

	hash := crypto.Keccak256Hash(
		common.HexToAddress(operatorAddress).Bytes(),
		common.LeftPadBytes(amount.Bytes(), 32),
		common.HexToAddress(tokenAddress).Bytes(),
		common.HexToAddress(toAddress).Bytes(),
		[]byte(nonce),
		common.RightPadBytes([]byte(PaymentPay), 32), // ethers.js 等价 utils.formatBytes32String('payment_pay')
	)
	hashByte := s.sha3Hash(hash.Bytes())
	hashStr = hexutil.Encode(hashByte)
	signature, err = s.sign(hashByte)

	fmt.Println("hashStr", hashStr)
	fmt.Println("signature", signature)
	fmt.Println("err", err)

}
