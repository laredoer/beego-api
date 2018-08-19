package util

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)


// 对字符串进行哈希加密
func Sha1(input string) string {
	if input == "" {
		return "adc83b19e793491b1c6ea0fd8b46cd9f32e592fc"
	}
	return fmt.Sprintf("%x", sha1.Sum([]byte(input)))
}

func Secret2Password(username, secret string) string {
	return Sha1(Sha1(secret[:8]) + Sha1(username) + Sha1(secret[8:]))
}


// 对字符串base64加密
func Base64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}
