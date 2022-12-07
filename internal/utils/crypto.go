package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

const (
	AesKey = "4ngxwh0z4lntrdwz"
	AesIV  = "lbg68ewassx86u8l"
)

// AesEny CTR模式
func AesEny(plaintext []byte) []byte {
	var (
		block cipher.Block
		err   error
	)
	if block, err = aes.NewCipher([]byte(AesKey)); err != nil {
		log.Fatal(err)
	}
	stream := cipher.NewCTR(block, []byte(AesIV))
	stream.XORKeyStream(plaintext, plaintext)
	return plaintext
}

// Base64AesEny 将加密 / 解密结果转换为base64
func Base64AesEny(plaintext []byte) string {
	return base64.StdEncoding.EncodeToString(AesEny(plaintext))
}

func Base64AesEnyDecode(text string) string {
	t, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}

	return string(AesEny(t))
}
