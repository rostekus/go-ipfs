package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomKey() string {
	keyLength := 8
	keyBytes := make([]byte, keyLength)

	_, err := rand.Read(keyBytes)
	if err != nil {
		panic(err)
	}

	keyName := hex.EncodeToString(keyBytes)

	return keyName
}
