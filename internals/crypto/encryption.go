package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
	"log"
)

func deriveKey(passphrase string) []byte {
	// Gera uma chave de 32 bytes usando SHA-256
	hash := sha256.Sum256([]byte(passphrase))
	return hash[:]
}

func Encrypt(data []byte, passphrase string) []byte {
	key := deriveKey(passphrase)

	aesC, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(aesC)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	return gcm.Seal(nonce, nonce, data, nil)
}
