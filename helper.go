package rocketfuel

import (
	"encoding/json"
	"fmt"

	"github.com/mervick/aes-everywhere/go/aes256"
)

func encrypt(toEncrypt string, secret string) string {
	// encryption
	encrypted := aes256.Encrypt(toEncrypt, secret)
	fmt.Println(encrypted, "encrypted")

	return encrypted
}
func marshalize(value any) []byte {
	toEncrypt, _ := json.Marshal(value)
	return toEncrypt
}
