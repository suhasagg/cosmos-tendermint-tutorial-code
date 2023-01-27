package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/cosmos/cosmos-sdk/crypto/commitment"
)

// struct to represent an IBC packet
type IBCPacket struct {
	SourceChain string `json:"source_chain"`
	DestChain   string `json:"dest_chain"`
	Data        []byte `json:"data"`
	Commitment  []byte `json:"commitment"`
}

// struct to represent a batch of IBC packets
type IBCBatch struct {
	Packets []IBCPacket `json:"packets"`
	Commitment  []byte `json:"commitment"`
}

// function to handle incoming IBC batches
func handleIBCBatch(w http.ResponseWriter, r *http.Request) {
	// parse the incoming batch
	var batch IBCBatch
	err := json.NewDecoder(r.Body).Decode(&batch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Verify the integrity of the packets using the batch commitment
	batchCommitment := commitment.NewCommitment(batch.Commitment)
	for _, packet := range batch.Packets {
		packetCommitment := commitment.NewCommitment(packet.Commitment)
		if !batchCommitment.Verify([]commitment.Commitment{packetCommitment}) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// process the packets
	for _, packet := range batch.Packets {
		fmt.Printf("Received IBC packet from %s to %s: %s\n", packet.SourceChain, packet.DestChain, string(packet.Data))
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	// set up an HTTP server to handle incoming IBC batches
	http.HandleFunc("/ibc", handleIBCBatch)
	http.ListenAndServe(":8000", nil)
}
