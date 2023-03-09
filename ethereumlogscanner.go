package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to Ethereum node
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/your-project-id")
	if err != nil {
		log.Fatal(err)
	}

	// Create bloom filter
	bloom := ethereum.NewBloomFilter()

	// Add event signatures to bloom filter
	bloom.Add([]byte("Transfer(address,address,uint256)"))
	bloom.Add([]byte("Approval(address,address,uint256)"))

	// Subscribe to logs
	logsCh := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), ethereum.FilterQuery{
		Topics: [][]common.Hash{{}},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Process logs
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case log := <-logsCh:
			// Check if bloom filter matches
			if bloom.Test(log.Topics) {
				// Decode event data
				var event struct {
					From  common.Address
					To    common.Address
					Value *big.Int
				}
				err := client.UnmarshalLog(&event, "Transfer", log)
				if err != nil {
					log.Println(err)
				} else {
					// Print event data
					data, _ := json.Marshal(event)
					fmt.Println(string(data))
				}
			}
		}
	}
}
