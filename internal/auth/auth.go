package auth

import (
	"math/rand"
)

func randomCode() string {
	return randomString(randomBetween5And20())
}

func randomBetween5And20() int {
	return rand.Intn(16) + 5 // rand.Intn(16) gives a number between 0 and 15, so add 5
}

func randomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
