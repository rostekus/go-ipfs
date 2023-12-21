package db

import (
	"database/sql"
	"log"
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

type DBAnalyzer struct {
	db *sql.DB
}

func New(uri string) *DBAnalyzer {
	db, err := sql.Open("mysql", uri)
	if err != nil {
		log.Fatal(err)
	}

	return &DBAnalyzer{db: db}
}

func (db *DBAnalyzer) Close() {
	if db.db != nil {
		db.db.Close()
	}
}

func (db *DBAnalyzer) GetTableinfo(tableName string) *TableInfo {
	columns, err := db.getColumns(tableName)
	if err != nil {
		log.Fatal(err)
	}

	// Get rows data
	rowsData, err := db.getRows(tableName, columns)
	if err != nil {
		log.Fatal(err)
	}

	// Create TableInfo struct
	tableInfo := TableInfo{
		TableName: tableName,
		Columns:   columns,
		Rows:      rowsData,
	}

	return &tableInfo
	// Convert to JSON
}

func (db *DBAnalyzer) getColumns(tableName string) ([]Column, error) {
	rows, err := db.db.Query("DESCRIBE " + tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []Column

	for rows.Next() {
		var columnName, columnType, nullable, key, defaultValue, extra sql.NullString

		if err := rows.Scan(&columnName, &columnType, &nullable, &key, &defaultValue, &extra); err != nil {
			return nil, err
		}

		column := Column{
			Name: columnName.String,
			Type: columnType.String,
			// You might want to include other column properties like nullable, key, default, extra here
		}

		columns = append(columns, column)
	}

	return columns, nil
}

func (db *DBAnalyzer) getRows(tableName string, columns []Column) ([]Row, error) {
	rows, err := db.db.Query("SELECT * FROM " + tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rowsData []Row

	for rows.Next() {
		row := make(Row)

		// Create a slice of interface{} to hold the row values
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		for i, col := range columns {
			val := values[i]
			row[col.Name] = string(val.([]uint8))
		}

		rowsData = append(rowsData, row)
	}

	return rowsData, nil
}
