package main

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomBytes(length int) ([]byte, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Create a byte slice to store the random bytes
	randomBytes := make([]byte, length)

	maxIndex := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return []byte{}, err
		}

		randomBytes[i] = charset[randomIndex.Int64()]
	}

	return randomBytes, nil
}
