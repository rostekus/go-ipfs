package db

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
