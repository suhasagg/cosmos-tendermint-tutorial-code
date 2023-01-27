package main

import (
	"github.com/tendermint/tendermint/abci/server"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/types"
)

func main() {
	// Create a new Tendermint node
	config := node.DefaultConfig()
	config.P2P.AddrBookStrict = false
	config.P2P.AllowDuplicateIP = true
	config.Snapshots.Interval = 60 // Snapshot sync every 60 blocks
	config.Snapshots.KeepRecent = 10 // Keep the 10 most recent snapshots
	config.Snapshots.Prune = true // Prune old snapshots
	
	// Create a new Tendermint node
	n, err := node.NewNode(config, types.LoadOrGenNodeKey(), p2p.NewDefaultNodeServer(), &server.Server{})
	if err != nil {
		log.Error(err.Error())
		return
	}

	// Start the node
	if err := n.Start(); err != nil {
		log.Error(err.Error())
		return
	}

	// Wait for the node to stop
	n.Wait()
}
