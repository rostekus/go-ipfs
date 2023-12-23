package ipfs

import (
	"encoding/json"
	"fmt"
)

type CIDResponse struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

func GetCIDResponse(jsonBody string) (*CIDResponse, error) {
	var cidBody CIDResponse

	// Unmarshal JSON data into the FileDetails struct
	err := json.Unmarshal([]byte(jsonBody), &cidBody)
	if err != nil {
		fmt.Println("Error:", err)
		return &CIDResponse{}, err
	}
	return &cidBody, nil
}
