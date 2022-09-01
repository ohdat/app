package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type AES struct {
	AesKey []byte // 对称秘钥长度必须是16的倍数
}

func New(key string) *AES {
	// key 要求必须是32位
	if len(key) != 32 {
		panic("key length must be 32")
	}
	return &AES{
		AesKey: []byte(key),
	}
}
func (a AES) Encode(key string) (s string, err error) {
	var b []byte
	b, err = a.AesEncrypt([]byte(key), a.AesKey)
	if err != nil {
		return
	}
	s = base64.StdEncoding.EncodeToString(b)
	return
}
func (a AES) Decode(s string) (str string, err error) {
	var (
		key []byte
		b   []byte
	)
	key, err = base64.StdEncoding.DecodeString(s)
	if err != nil {
		return
	}
	b, err = a.AesDecrypt(key, a.AesKey)
	if err != nil {
		return
	}
	return string(b), nil
}

func (a AES) PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (a AES) PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//AES加密,CBC
func (a AES) AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = a.PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AES解密
func (a AES) AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = a.PKCS7UnPadding(origData)
	return origData, nil
}
