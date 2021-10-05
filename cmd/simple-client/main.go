package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ETH_NODE_URL = "http://127.0.0.1:8545"
)

// Connecting to
func main() {

	client, err := ethclient.Dial(ETH_NODE_URL)

	if err != nil {
		log.Fatal("Oops! There was a problem", err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}
	defer client.Close()

	header, err := client.HeaderByNumber(context.Background(), nil)
	// id, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Cannot get the chain id", err)
	} else {
		fmt.Printf("hash id %d\n", header)
	}

}
