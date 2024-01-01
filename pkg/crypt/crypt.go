package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

func EnsureCipherKeyInFile(fileName string) string {
	// Read the file content
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("Creating %s", fileName)
		err = os.WriteFile(fileName, []byte("CIPHER_KEY="), 0644)
		if err != nil {
			log.Fatal(err)
		}
		return EnsureCipherKeyInFile(fileName)
	}

	// Check if the key exists and is not empty
	var key string
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "CIPHER_KEY=") {
			key = strings.TrimPrefix(line, "CIPHER_KEY=")
			break
		}
	}

	// If key is empty, generate a new one and update the file
	if key == "" {
		log.Println("Generating new CIPHER_KEY")
		key = GenerateRandom32ByteKey()
		newContent := strings.Replace(string(content), "CIPHER_KEY=", "CIPHER_KEY="+key, 1)
		err = os.WriteFile(fileName, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	return key
}

// Returns a 32 byte string key (16 byte hexadecimal key)
func GenerateRandom32ByteKey() string {
	key, err := GenerateRandomString(16)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

// Returns a 64 byte string key (32 byte hexadecimal key)
func GenerateRandom64ByteKey() string {
	key, err := GenerateRandomString(32)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

func GenerateRandomString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func Encrypt(data []byte, passphrase string) (string, error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(encrypted string, passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	data, err := hex.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
