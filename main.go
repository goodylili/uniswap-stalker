package main

import (
	v2 "UniswapStalker/v2"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

const (
	infuraURL        = "wss://g.w.lavanet.xyz:443/gateway/eth/rpc/a14c2dda303d3522c6f553a37b187f63"
	uniswapv2Factory = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
)

func main() {
	// Connect to the Ethereum client
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()

	// Replace with the actual address of the Uniswap V2 factory contract
	factoryAddress := common.HexToAddress(uniswapv2Factory)

	// Create an instance of the factory contract
	factory, err := v2.NewV2Factory(factoryAddress, client)
	if err != nil {
		log.Fatalf("Failed to create instance of factory contract: %v", err)
	}

	// Set up a filter query to listen for PairCreated events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{factoryAddress},
	}

	// Subscribe to the PairCreated events
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Listening for PairCreated events...")

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Error while listening for logs: %v", err)
		case vLog := <-logs:
			// Debugging: print the raw log
			fmt.Printf("Raw log: %+v\n", vLog)

			// Parse the log data
			event, err := factory.ParsePairCreated(vLog)
			if err != nil {
				log.Fatalf("Failed to parse log: %v", err)
			}

			// Log the event details
			fmt.Printf("New Pair Created: \nToken0: %s\nToken1: %s\nPair Address: %s\n", event.Token0.Hex(), event.Token1.Hex(), event.Pair.Hex())
		}
	}
}
