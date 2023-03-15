package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type AES struct {
	iv    []byte
	block cipher.Block
}

func New(key string) *AES {
	//The length of the symmetric key must be a multiple of 16.
	if len(key)%16 != 0 {
		panic("key length must be 16 or 32")
	}
	// NewCipher creates and returns a new cipher.Block.
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	return &AES{
		iv:    []byte(key)[:block.BlockSize()],
		block: block,
	}
}

// pkcs7Padding
func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("data is empty")
	}
	// 获取填充的个数
	// The last byte of the data is the number of padding bytes.
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// Encrypt Encrypt
func (s *AES) Encrypt(data []byte) ([]byte, error) {
	//get the block size
	blockSize := s.block.BlockSize()
	// padding
	// The length of the data must be a multiple of the block size.
	encryptBytes := pkcs7Padding(data, blockSize)
	// Create a new byte array the size of the crypted text
	crypted := make([]byte, len(encryptBytes))
	//using cbc encrypt mode
	blockMode := cipher.NewCBCEncrypter(s.block, s.iv)
	// CryptBlocks can work in-place if the two arguments are the same.
	blockMode.CryptBlocks(crypted, encryptBytes)
	return crypted, nil
}

// Decrypt Decrypt
func (s *AES) Decrypt(data []byte) ([]byte, error) {
	var err error
	//using cbc decrypt mode
	blockMode := cipher.NewCBCDecrypter(s.block, s.iv)
	// Create a new byte array the size of the original data
	crypted := make([]byte, len(data))
	// CryptBlocks can work in-place if the two arguments are the same.
	blockMode.CryptBlocks(crypted, data)
	// Remove the padding
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return nil, err
	}
	return crypted, nil
}

func (s *AES) Encode(key string) (str string, err error) {
	var b []byte
	b, err = s.Encrypt([]byte(key))
	if err != nil {
		return
	}
	str = base64.StdEncoding.EncodeToString(b)
	return
}
func (s *AES) Decode(message string) (str string, err error) {
	var (
		key []byte
		b   []byte
	)
	key, err = base64.StdEncoding.DecodeString(message)
	if err != nil {
		return
	}
	b, err = s.Decrypt(key)
	if err != nil {
		return
	}
	return string(b), nil
}
