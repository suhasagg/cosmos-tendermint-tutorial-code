package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinebox/graphql"
)

type Block struct {
	Hash         string `json:"hash"`
	Number       uint64 `json:"number"`
	Transactions []struct {
		Hash   string `json:"hash"`
		From   string `json:"from"`
		To     string `json:"to"`
		Value  string `json:"value"`
		Input  string `json:"input"`
		Gas    uint64 `json:"gas"`
		GasFee string `json:"gasFee"`
	} `json:"transactions"`
}

type Response struct {
	Block Block `json:"block"`
}

func main() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}

	gqlClient := graphql.NewClient("http://localhost:8547/graphql")

	req := graphql.NewRequest(`
		query ($blockNumber: Int!) {
			block(number: $blockNumber) {
				hash
				number
				transactions {
					hash
					from
					to
					value
					input
					gas
					gasFee
				}
			}
		}
	`)

	req.Var("blockNumber", 123456)

	var respData Response
	if err := gqlClient.Run(context.Background(), req, &respData); err != nil {
		panic(err)
	}

	blockHash := hexutil.Encode([]byte(respData.Block.Hash))
	blockNumber := respData.Block.Number

	block, err := client.BlockByHash(context.Background(), common.HexToHash(blockHash))
	if err != nil {
		panic(err)
	}

	fmt.Println("Block Number:", blockNumber)
	fmt.Println("Block Hash:", block.Hash().Hex())
	fmt.Println("Block Time:", block.Time())
	fmt.Println("Number of Transactions:", len(block.Transactions()))
}
