package wallet

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type PaymentPaySignKey string

const (
	PaymentRedeem PaymentPaySignKey = "payment_redeem"
	PaymentPay    PaymentPaySignKey = "payment_pay"
)

// PaymentPay payment pay 的签名
func (s Signature) PaymentPay(
	ownerAddress string,
	amount big.Int,
	tokenAddress,
	toAddress string,
	nonce string,
	key string,
) (hashStr, signature string, err error) {
	//hexStr := "iq20230213072752447521620194513339"
	//utfBytes := string.ToUtf8Bytes(nonce)
	//string.ToUtf8Bytes(nonce)
	_, err = common.ParseHexOrString(nonce)
	if err != nil {
		return
	}
	hash := crypto.Keccak256Hash(
		common.HexToAddress(ownerAddress).Bytes(),
		common.LeftPadBytes(amount.Bytes(), 32),
		common.HexToAddress(tokenAddress).Bytes(),
		common.HexToAddress(toAddress).Bytes(),
		[]byte(nonce),
		common.RightPadBytes([]byte(key), 32), // ethers.js 等价 utils.formatBytes32String('payment_pay')
	)

	hashByte := s.sha3Hash(hash.Bytes())
	hashStr = hexutil.Encode(hashByte)
	signature, err = s.sign(hashByte)
	return
}

// PaymentRedeem payment redeem 的签名
func (s Signature) PaymentRedeem(
	ownerAddress string,
	amount big.Int,
	tokenAddress string,
	fromAddress string,
	blockNum int,
	nonce string,
	key string,
) (hashStr, signature string, err error) {

	hash := crypto.Keccak256Hash(
		common.HexToAddress(ownerAddress).Bytes(),
		common.LeftPadBytes(amount.Bytes(), 32),
		common.HexToAddress(tokenAddress).Bytes(),
		common.HexToAddress(fromAddress).Bytes(),
		common.LeftPadBytes(big.NewInt(int64(blockNum)).Bytes(), 32),
		[]byte(nonce),
		[]byte(key),
	)
	hashByte := s.sha3Hash(hash.Bytes())
	hashStr = hexutil.Encode(hashByte)
	signature, err = s.sign(hashByte)
	return
}
