package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(bytes []byte) string {
	hash := md5.Sum(bytes)
	return hex.EncodeToString(hash[:])
}
