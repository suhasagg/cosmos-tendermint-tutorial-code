package main

import (
    "fmt"
)

type Transaction struct {
    ID      int
    Amount  float64
}

func main() {
    mempool := make(chan Transaction, 10)
    //won't wait for transaction processing for tx ingestion

    // start a goroutine to process transactions
    go func() {
        for {
            tx := <-mempool // read from the channel
            fmt.Printf("Processing transaction %d for amount %.2f\n", tx.ID, tx.Amount)
            // process the transaction
        }
    }()

    // add some transactions to the mempool
    for i := 1; i <= 5; i++ {
        tx := Transaction{ID: i, Amount: float64(i * 100)}
        mempool <- tx // send to the channel
    }

    // wait for the transactions to be processed
    for i := 1; i <= 5; i++ {
        <-mempool // read from the channel to block until processing is done
    }
}
