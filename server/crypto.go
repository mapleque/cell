package server

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
)

// EncryptAesEcbBase64 EncryptAesEcbBase64 encrypt data with key using AES-ECB
// and encode with base64.
//
// The key should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256
func EncryptAesEcbBase64(data []byte, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	src := _PKCS5Padding(data, blockSize)
	dst := make([]byte, 0)
	tmp := make([]byte, block.BlockSize())
	for len(src) > 0 {
		block.Encrypt(tmp, src[:blockSize])
		src = src[blockSize:]
		dst = append(dst, tmp...)
	}
	return string(base64.StdEncoding.EncodeToString(dst)), nil
}

// DecryptAesEcbBase64 DecryptAesEcbBase64 decode data with base64
// and decrypt with key using AES-ECB.
//
// The key should be the AES key,
// either 16, 24, or 32 bytes to select
// AES-128, AES-192, or AES-256
func DecryptAesEcbBase64(data, key string) ([]byte, error) {
	src, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, errEncDataInvalid
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	dst := make([]byte, 0)
	tmp := make([]byte, blockSize)
	if len(src)%blockSize != 0 {
		return nil, errEncDataInvalid
	}
	for len(src) > 0 {
		block.Decrypt(tmp, src[:blockSize])
		src = src[blockSize:]
		dst = append(dst, tmp...)
	}
	res := _PKCS5UnPadding(dst)
	return res, nil
}

func _PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func _PKCS5UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	datalen := length - unpadding
	if datalen < 0 || datalen > len(data) {
		return nil
	}
	return data[:datalen]
}
