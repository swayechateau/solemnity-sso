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

	"golang.org/x/crypto/sha3"
)

// EnsureCipherKeyInFile ensures that a cipher key is present in the specified file.
// If the file does not exist, it creates the file with an empty cipher key.
// If the file exists but the cipher key is empty, it generates a new cipher key and updates the file.
// If the file exists and the cipher key is already present, it returns the existing cipher key.
// The function returns the cipher key as a string.
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

// GenerateRandom32ByteKey generates a random 32-byte key.
// It calls the GenerateRandomString function to generate a random string of length 16,
// and returns the generated key.
func GenerateRandom32ByteKey() string {
	key, err := GenerateRandomString(16)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

// GenerateRandom64ByteKey generates a random 64-byte key.
// It uses the GenerateRandomString function to generate a 32-byte random string,
// and returns the generated key.
func GenerateRandom64ByteKey() string {
	key, err := GenerateRandomString(32)
	if err != nil {
		log.Fatal(err)
	}
	return key
}

// GenerateRandomString generates a random string of length n.
// It uses cryptographic random number generator to ensure randomness.
// The generated string is returned in hexadecimal format.
// If an error occurs during the generation process, an error is returned.
func GenerateRandomString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// Encrypt encrypts the given data using the provided passphrase.
// It returns the encrypted data as a hexadecimal string and any error encountered.
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

// Decrypt decrypts the given encrypted string using the provided passphrase.
// It returns the decrypted data as a byte slice.
// If an error occurs during decryption, it returns nil and the corresponding error.
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

// HashStringToString takes a string as input and returns its SHA3-256 hash as a hexadecimal string.
func HashStringToString(data string) string {
	hashArray := sha3.Sum256([]byte(data))
	return hex.EncodeToString(hashArray[:])
}

// HashString calculates the SHA3-256 hash of the input string.
// It takes a string as input and returns a fixed-size array of 32 bytes.
func HashString(data string) [32]byte {
	return sha3.Sum256([]byte(data))
}

// HashByte calculates the SHA3-256 hash of the given byte slice.
// It returns a fixed-size array of 32 bytes representing the hash.
func HashByte(data []byte) [32]byte {
	return sha3.Sum256([]byte(data))
}
