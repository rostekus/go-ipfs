package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/rostekus/ipfs-go/internal/db"
)

func GenerateChainHash(rows []db.Row) string {
	var hashChain []string
	var previousHash string

	for _, row := range rows {
		rowData, err := json.Marshal(row)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			continue
		}

		dataToHash := string(rowData) + previousHash

		hash := sha256.Sum256([]byte(dataToHash))
		hashString := fmt.Sprintf("%x", hash)

		hashChain = append(hashChain, hashString)

		previousHash = hashString
	}

	finalHash := hashChain[len(hashChain)-1]
	return finalHash
}
