package utils

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	DBURI       string
	IpfsAddress string
	TableName   string
}

func GetArgs() *Args {
	dbFlag := flag.String("db", "", "MySQL database connection string")
	ipfsFlag := flag.String("ipfs", "", "IPFS address")

	tableFlag := flag.String("table", "", "Table name")
	flag.Parse()

	dbConnectionString := *dbFlag
	ipfsAddress := *ipfsFlag
	tableName := *tableFlag
	if dbConnectionString == "" || ipfsAddress == "" {
		fmt.Println("Usage: go run main.go --db=<db_connection_string> --ipfs=<ipfs_address>")
		os.Exit(-1)
		return &Args{}
	}
	args := &Args{
		DBURI: dbConnectionString, IpfsAddress: ipfsAddress,
		TableName: tableName,
	}
	return args
}
