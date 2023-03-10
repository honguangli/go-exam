package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
