package benchmarks

import cp "github.com/gtank/cryptopasta"

func EncryptCryptopasta(plaintext []byte) []byte {
	ciphertext, _ := cp.Encrypt(plaintext, cp.NewEncryptionKey())
	return ciphertext
}
