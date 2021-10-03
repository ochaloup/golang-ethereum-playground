package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	ETH_NODE_URL = "http://127.0.0.1:30303"
)

func main() {

	client, err := ethclient.Dial(ETH_NODE_URL)

	if err != nil {
		log.Fatal("Oops! There was a problem", err)
	} else {
		fmt.Println("Success! you are connected to the Ethereum Network")
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(header.Number.String())
	}
}
