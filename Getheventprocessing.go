package main

import (
    "context"
    "fmt"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    // Connect to a local Geth node
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        panic(err)
    }

    // Subscribe to new block headers
    headerCh := make(chan *types.Header)
    sub, err := client.SubscribeNewHead(context.Background(), headerCh)
    if err != nil {
        panic(err)
    }

    // Process new block headers
    go func() {
        for {
            select {
            case err := <-sub.Err():
                panic(err)
            case header := <-headerCh:
                fmt.Printf("New block header: %s\n", header.Hash().Hex())
            }
        }
    }()

    // Subscribe to new transactions
    txCh := make(chan *types.Transaction)
    sub, err = client.SubscribeNewTransactions(context.Background(), txCh)
    if err != nil {
        panic(err)
    }

    // Process new transactions
    go func() {
        for {
            select {
            case err := <-sub.Err():
                panic(err)
            case tx := <-txCh:
                fmt.Printf("New transaction: %s\n", tx.Hash().Hex())
            }
        }
    }()

    // Subscribe to contract events
    contractAddr := common.HexToAddress("0x1234567890123456789012345678901234567890")
    eventCh := make(chan *types.Log)
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddr},
    }
    sub, err = client.SubscribeFilterLogs(context.Background(), query, eventCh)
    if err != nil {
        panic(err)
    }

    // Process contract events
    go func() {
        for {
            select {
            case err := <-sub.Err():
                panic(err)
            case event := <-eventCh:
                fmt.Printf("New contract event: %s\n", event.Topics[0].Hex())
                // Process the event data here
                // For example, decode the event data and log the values
                var data MyEventData
                err := contractAbi.UnpackIntoInterface(&data, "MyEvent", event.Data)
                if err != nil {
                    panic(err)
                }
                fmt.Printf("Event values: %v\n", data)
            }
        }
    }()

    // Wait for events
    select {}
}

type MyEventData struct {
    Arg1 string
    Arg2 uint64
    Arg3 bool
}
