package main

import (
    "fmt"
    "sync"
    "time"
)

type Mempool struct {
    transactions map[string]*Transaction
    lock sync.RWMutex
    eviction time.Duration
}

type Transaction struct {
    ID        string
    Data      []byte
    Timestamp time.Time
}

func (m *Mempool) AddTransaction(tx Transaction) {
    m.lock.Lock()
    defer m.lock.Unlock()

    m.transactions[tx.ID] = &tx
}

func (m *Mempool) GetTransactions() []Transaction {
    m.lock.RLock()
    defer m.lock.RUnlock()

    var transactions []Transaction
    for _, tx := range m.transactions {
        transactions = append(transactions, *tx)
    }

    return transactions
}

func (m *Mempool) EvictExpiredTransactions() {
    m.lock.Lock()
    defer m.lock.Unlock()

    now := time.Now()
    for id, tx := range m.transactions {
        if now.Sub(tx.Timestamp) > m.eviction {
            delete(m.transactions, id)
        }
    }
}

func main() {
    mempool := &Mempool{
        transactions: make(map[string]*Transaction),
        eviction: 10 * time.Second,
    }

    tx1 := Transaction{ID: "1", Data: []byte("Hello World"), Timestamp: time.Now()}
    tx2 := Transaction{ID: "2", Data: []byte("Tendermint Rocks!"), Timestamp: time.Now()}

    mempool.AddTransaction(tx1)
    mempool.AddTransaction(tx2)

    // Wait for eviction time
    time.Sleep(11 * time.Second)
    mempool.EvictExpiredTransactions()

    transactions := mempool.GetTransactions()
    for _, tx := range transactions {
        fmt.Printf("Transaction ID: %s, Data: %s\n", tx.ID, string(tx.Data))
    }
    
}
