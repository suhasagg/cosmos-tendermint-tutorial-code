package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    n = 4 // number of nodes in the system
    f = 1 // number of faults tolerated
)

type Node struct {
    ID    int
    Round int
    Voted bool
    Vote  int
}

type PBFT struct {
    nodes []Node
    mu    sync.RWMutex
}

func NewPBFT() *PBFT {
    nodes := make([]Node, n)
    for i := 0; i < n; i++ {
        nodes[i] = Node{ID: i}
    }
    return &PBFT{nodes: nodes}
}

func (p *PBFT) Broadcast(round int, value int) {
    p.mu.Lock()
    defer p.mu.Unlock()
    for i := 0; i < n; i++ {
        p.nodes[i].Round = round
        if !p.nodes[i].Voted {
            p.nodes[i].Vote = value
            p.nodes[i].Voted = true
        }
    }
}

func (p *PBFT) CheckConsensus(round int) (bool, int) {
    p.mu.RLock()
    defer p.mu.RUnlock()
    votes := make(map[int]int)
    for i := 0; i < n; i++ {
        if p.nodes[i].Round == round {
            votes[p.nodes[i].Vote]++
        }
    }
    maxVotes := 0
    var maxValue int
    for value, count := range votes {
        if count > maxVotes {
            maxVotes = count
            maxValue = value
        }
    }
    return maxVotes > n-f, maxValue
}

func main() {
    pbft := NewPBFT()

    // simulate client request
    round := 1
    value := rand.Intn(100)
    fmt.Println("Client request: value =", value)

    // broadcast request to nodes
    pbft.Broadcast(round, value)

    // wait for consensus
    time.Sleep(time.Second)

    // check for consensus
    consensus, agreedValue := pbft.CheckConsensus(round)
    if consensus {
        fmt.Println("Consensus reached: value =", agreedValue)
    } else {
        fmt.Println("No consensus reached")
    }
}
