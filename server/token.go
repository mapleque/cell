package server

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

// RandToken RandToken general a 64 bytes random string.
//
// The token combine with 2 random md5 (64 bytes),
// which charactors order has been rebuild as a new strange mode.
func RandToken() string {
	return RandMd5() + RandMd5()
}

func RandMd5() string {
	return MD5(fmt.Sprintf("%d%d%d", rand.Intn(10000), time.Now().Unix(), rand.Intn(10000)))
}

func MD5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	data := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(data)))
	hex.Encode(dst, data)
	return string(dst)
}
