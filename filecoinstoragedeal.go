package main

import (
    "fmt"
    "log"
    "github.com/filecoin-project/go-address"
    "github.com/filecoin-project/go-state-types/abi"
    "github.com/filecoin-project/go-state-types/big"
    "github.com/filecoin-project/go-filecoin/actor"
    "github.com/filecoin-project/go-filecoin/protocol/storage/storagedeal"
    "github.com/filecoin-project/go-filecoin/types"
)

func main() {
    // Define the client's address
    clientAddr, err := address.NewFromString("t3j4zfh6u4hwwjl2w4u4t4z4j5h5z5x5y5z5f5j5g5")
    if err != nil {
        log.Fatal(err)
    }

    // Define the provider's address
    providerAddr, err := address.NewFromString("t3j4zfh6u4hwwjl2w4u4t4z4j5h5z5x5y5z5f5j5g6")
    if err != nil {
        log.Fatal(err)
    }

    // Define the storage deal parameters
    dealParams := &storagedeal.Params{
        DataRef: &types.DataRef{
            TransferType: types.TTGraphsync,
            Root:         types.Cid{},
        },
        EpochPrice:  big.NewInt(1),
        MinPieceSize: abi.PaddedPieceSize(1 << 20),
    }

    // Propose the storage deal
    proposalCid, err := actor.ProposeStorageDeal(clientAddr, providerAddr, dealParams)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Storage deal proposed with CID: ", proposalCid)
}
