package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	shell "github.com/ipfs/go-ipfs-api"
	"github.com/rostekus/ipfs-go/internal/db"
	"github.com/rostekus/ipfs-go/internal/ipfs"
	"github.com/rostekus/ipfs-go/internal/utils"
)

func main() {
	args := utils.GetArgs()

	fmt.Printf("analazying table %s\n", args.TableName)
	dbAnalyzer := db.New(args.DBURI)
	defer dbAnalyzer.Close()

	tableInfo := dbAnalyzer.GetTableinfo(args.TableName)
	fmt.Printf("Table: %s\nColumns: %s\n", tableInfo.TableName, tableInfo.Columns)

	fmt.Println("Calculating Hash...")
	hashChain := utils.GenerateChainHash(tableInfo.Rows)
	fmt.Printf("Hash Chain: %s\n", hashChain)

	fmt.Printf("Obtaining CID, IPFS %s\n", args.IpfsAddress)
	client := ipfs.Client{Url: args.IpfsAddress}
	resp, err := client.Add(hashChain)
	if err != nil {
		log.Fatalf("couldn't obtained cid %s", err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("cannot read resp from ipfs %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("couldn't obtained cid, response status %d", resp.StatusCode)
	}
	cidBody, err := ipfs.GetCIDResponse(string(body))
	if err != nil {
		log.Fatalf("cannot read resp from ipfs %s", err.Error())
	}
	fmt.Printf("Obtained CID %s\n", cidBody.Hash)

	// ======= Using 3rd party library

	keyS := utils.GenerateRandomKey()
	fmt.Println("your key", keyS)
	sh := shell.NewShell("localhost:5001")
	key, err := sh.KeyGen(context.Background(), keyS)
	if err != nil {
		fmt.Println("Error creating key:", err)
		return
	}

	// Publish the CID to IPNS with the created key
	ipnsHash, err := sh.PublishWithDetails(cidBody.Hash, key.Id, 0, 0, false)
	if err != nil {
		fmt.Println("Error publishing to IPNS:", err)
		return
	}

	fmt.Printf("IPNS: %s of the content with cid: %s\n", ipnsHash.Name, ipnsHash.Value)
}
