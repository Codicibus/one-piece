package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}