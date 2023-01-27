package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	// broadcastInterval is the interval at which blocks are broadcasted to the network
	broadcastInterval = 5 * time.Second
)

var (
	// blockChan is the channel used to broadcast blocks
	blockChan = make(chan []byte)
	// knownNodes is a map of known nodes in the network
	knownNodes = make(map[string]bool)
	// knownNodesLock is a lock for accessing the knownNodes map
	knownNodesLock sync.RWMutex
)

func main() {
	// start the block broadcast routine
	go broadcastRoutine()

	// start the listener routine
	go listenRoutine()

	// add some known nodes to the network
	addKnownNode("127.0.0.1:8000")
	addKnownNode("127.0.0.1:8001")

	// simulate a new block being mined
	block := []byte("Block 1")
	blockChan <- block

	// wait for the program to exit
	select {}
}

// broadcastRoutine broadcasts blocks to the network at a fixed interval
func broadcastRoutine() {
	for {
		select {
		case block := <-blockChan:
			broadcastBlock(block)
		case <-time.After(broadcastInterval):
			// broadcast any blocks that may have been missed
			for len(blockChan) > 0 {
				block := <-blockChan
				broadcastBlock(block)
			}
		}
	}
}

// broadcastBlock sends a block to all known nodes in the network
func broadcastBlock(block []byte) {
	knownNodesLock.RLock()
	defer knownNodesLock.RUnlock()
	for node := range knownNodes {
		go sendBlock(node, block)
	}
}

// sendBlock sends a block to a specific node
func sendBlock(node string, block []byte) {
	conn, err := net.Dial("tcp", node)
	if err != nil {
		fmt.Println("Error connecting to node:", err)
		return
	}
	defer conn.Close()
	conn.Write(block)
}

// listenRoutine listens for incoming blocks from other nodes
func listenRoutine() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error starting listener:", err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
