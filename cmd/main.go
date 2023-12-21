package main

import (
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rostekus/ipfs-go/internal/db"
)

type TableInfo struct {
	TableName string   `json:"tableName"`
	Columns   []Column `json:"columns"`
	Rows      []Row    `json:"rows"`
}

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Row map[string]interface{}

func main() {
	// Connect to MySQL database
	uri := "root:examplepassword@tcp(localhost:3306)/exampledb"
	tableName := "users"

	dbAnalyzer := db.New(uri)
	defer dbAnalyzer.Close()

	// Create TableInfo struct
	tableInfo := dbAnalyzer.GetTableinfo(tableName)
	fmt.Print(tableInfo)

	// Convert to JSON
	jsonData, err := json.Marshal(tableInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
