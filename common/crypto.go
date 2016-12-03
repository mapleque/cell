package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"

	"github.com/coral/log"
)

func Md5(tar string) string {
	h := md5.New()
	h.Write([]byte(tar))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func AesEcbEnc(key, tar string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Error("key error", err.Error())
		return ""
	}
	blockSize := block.BlockSize()
	ecb := cipher.NewCBCEncrypter(block, []byte(key)[:blockSize])
	src := PKCS5Padding([]byte(tar), blockSize)
	ret := make([]byte, len(src))
	ecb.CryptBlocks(ret, src)
	res := base64.StdEncoding.EncodeToString(ret)
	return res
}

func AesEcbDec(key, tar string) string {
	src, err := base64.StdEncoding.DecodeString(tar)
	if err != nil {
		log.Error("tar error", err.Error())
		return ""
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Error("key error", err.Error())
		return ""
	}
	blockSize := block.BlockSize()
	ecb := cipher.NewCBCDecrypter(block, []byte(key)[:blockSize])

	ret := make([]byte, len(src))
	ecb.CryptBlocks(ret, src)

	res := PKCS5UnPadding(ret)

	return string(res)
}

func PKCS5Padding(tar []byte, blockSize int) []byte {
	padding := blockSize - len(tar)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(tar, padtext...)
}
func PKCS5UnPadding(tar []byte) []byte {
	length := len(tar)
	unpadding := int(tar[length-1])
	return tar[:(length - unpadding)]
}
