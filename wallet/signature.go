package wallet

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Signature struct {
	HexPrivateKey string
}

func NewSignature(hexPrivateKey string) Signature {
	//var hexPrivateKey = viper.GetString("sign_wallet.private_key")
	if len(hexPrivateKey) != 66 {
		panic(`privateKey: fail`)
	}
	return Signature{
		HexPrivateKey: hexPrivateKey,
	}
}

// sign 签名
func (s Signature) sign(data []byte) (signature string, err error) {
	var (
		privateKey    *ecdsa.PrivateKey
		signatureByte []byte
	)
	privateKey, err = crypto.HexToECDSA(s.HexPrivateKey[2:])
	if err != nil {
		return
	}
	signatureByte, err = crypto.Sign(data, privateKey)
	if err != nil {
		return
	}
	signatureByte[64] += 27
	signature = hexutil.Encode(signatureByte)
	return
}

// sha3Hash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//
//	keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func (s Signature) sha3Hash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func (s Signature) Sign(msg []byte) (string, error) {
	return s.sign(s.sha3Hash(msg))
}

// signTypedData eip712 sign
func (s Signature) SignTypedData(data []byte) (string, error) {
	var message apitypes.TypedData
	err := json.Unmarshal(data, &message)
	if err != nil {
		return "", err
	}
	msg, err := signV4Byte(message)
	if err != nil {
		return "", err
	}
	return s.sign(msg)
}
func fixMessage(typedData apitypes.TypedData) apitypes.TypedData {
	// opensea listing message has the extra data totalOriginalConsiderationItems
	// typedData.Message remove totalOriginalConsiderationItems field
	delete(typedData.Message, "totalOriginalConsiderationItems")
	return typedData
}
func signV4Byte(typedData apitypes.TypedData) ([]byte, error) {
	typedData = fixMessage(typedData)
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		fmt.Printf("typedDataHash failed '%v'", err)
		return nil, err
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sighash := crypto.Keccak256(rawData)
	return sighash, nil
}

func (s Signature) VerifySign(from, sigHex string, msg []byte) bool {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v", p)
			debug.PrintStack()
		}
	}()
	fromAddr := common.HexToAddress(from)
	sig := hexutil.MustDecode(sigHex)
	if sig[64] != 27 && sig[64] != 28 {
	} else {
		sig[64] -= 27
	}
	pubKey, err := crypto.SigToPub(s.sha3Hash(msg), sig)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return fromAddr == recoveredAddr
}

// SigToPub 验证签名 返回解析后的地址
func (s Signature) SigToPub(sigHex string, msg []byte) string {
	sig := hexutil.MustDecode(sigHex)
	if sig[64] != 27 && sig[64] != 28 {
	} else {
		sig[64] -= 27
	}
	pubKey, err := crypto.SigToPub(s.sha3Hash(msg), sig)
	if err != nil {
		log.Println("SigToPubErr:", err)
		return ""
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr.String()
}
