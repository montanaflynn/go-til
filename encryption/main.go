package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
)

func createHash(key string) []byte {
	h := sha256.New()
	h.Write([]byte(key))
	return h.Sum(nil)
}

// Encrypt encrypts data using the passphrase.
func Encrypt(data []byte, passphrase string) ([]byte, error) {
	hash := createHash(passphrase)
	block, err := aes.NewCipher(hash)
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Decrypt decrypts data using the passphrase.
func Decrypt(data []byte, passphrase string) ([]byte, error) {
	key := createHash(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return []byte{}, err
	}
	return plaintext, nil
}

func main() {
	fmt.Println("Starting the application...")
	ciphertext, err := Encrypt([]byte("Hello World"), "password")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted: %x\n", ciphertext)
	plaintext, err := Decrypt(ciphertext, "password")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", plaintext)
}
