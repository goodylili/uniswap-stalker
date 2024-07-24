package main

//
//import (
//	"context"
//	"fmt"
//	"log"
//	"math/big"
//	"time"
//
//	"github.com/ethereum/go-ethereum"
//	"github.com/ethereum/go-ethereum/accounts/abi/bind"
//	"github.com/ethereum/go-ethereum/common"
//	"github.com/ethereum/go-ethereum/core/types"
//	"github.com/ethereum/go-ethereum/ethclient"
//	"github.com/your_package/v2" // Replace with the actual import path of your generated factory package
//)
//
//func main() {
//	// Connect to the Ethereum client using WebSocket
//	client, err := ethclient.DialContext(context.Background(), "wss://mainnet.infura.io/ws/v3/YOUR_INFURA_PROJECT_ID")
//	if err != nil {
//		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
//	}
//
//	// Replace with the actual address of the Uniswap V2 factory contract
//	factoryAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
//
//	// Create an instance of the factory contract
//	factory, err := v2.NewV2Factory(factoryAddress, client)
//	if err != nil {
//		log.Fatalf("Failed to create instance of factory contract: %v", err)
//	}
//
//	// Set up a filter query to listen for PairCreated events
//	query := ethereum.FilterQuery{
//		Addresses: []common.Address{factoryAddress},
//	}
//
//	// Subscribe to the PairCreated events
//	logs := make(chan types.Log)
//	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
//	if err != nil {
//		log.Fatalf("Failed to subscribe to logs: %v", err)
//	}
//
//	fmt.Println("Listening for PairCreated events...")
//
//	for {
//		select {
//		case err := <-sub.Err():
//			log.Fatalf("Error while listening for logs: %v", err)
//		case vLog := <-logs:
//			// Parse the log data
//			event, err := factory.ParsePairCreated(vLog)
//			if err != nil {
//				log.Fatalf("Failed to parse log: %v", err)
//			}
//
//			// Log the event details
//			fmt.Printf("New Pair Created: \nToken0: %s\nToken1: %s\nPair Address: %s\n", event.Token0.Hex(), event.Token1.Hex(), event.Pair.Hex())
//		}
//	}
//}
