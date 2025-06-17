package benchmarks

import (
	"crypto/aes"
	"crypto/cipher"
)

var key = []byte("thisis32bitlongpassphraseimusing") // 32 bytes AES-256
var iv = []byte("thisis16bytesiv!")                  // 16 byte IV

func EncryptStandard(plaintext []byte) []byte {
	block, _ := aes.NewCipher(key)
	stream := cipher.NewCTR(block, iv)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}
