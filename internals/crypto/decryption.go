package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

func Decrypt(data []byte, passphrase string) []byte {
	key := deriveKey(passphrase)

	aesC, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(aesC)
	if err != nil {
		log.Fatal(err)
	}

	nonceSize := gcm.NonceSize()

	if len(data) < nonceSize {
		log.Fatal("WRONG!")
	}

	nonce, data := data[:nonceSize], data[nonceSize:]
	text, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		log.Fatal(err)
	}

	return text
}
