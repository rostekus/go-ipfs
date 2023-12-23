package db

import (
	"database/sql"
	"log"
)

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

	rowsData, err := db.getRows(tableName, columns)
	if err != nil {
		log.Fatal(err)
	}

	tableInfo := TableInfo{
		TableName: tableName,
		Columns:   columns,
		Rows:      rowsData,
	}

	return &tableInfo
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
