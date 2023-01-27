package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// create a new Ethereum client
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		fmt.Println("Error connecting to Ethereum client:", err)
		return
	}

	// define the transaction hash
	txHash := common.HexToHash("0x<transaction hash>")

	// get the transaction receipt
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		fmt.Println("Error getting transaction receipt:", err)
		return
	}

	// get the gas used by the transaction
	gasUsed := receipt.GasUsed

	// get the gas price of the transaction
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		fmt.Println("Error getting transaction:", err)
		return
	}

	if isPending {
		fmt.Println("Transaction is still pending.")
		return
	}

	gasPrice := tx.GasPrice()

	// calculate the gas cost of the transaction
	gasCost := new(big.Int).Mul(big.NewInt(int64(gasUsed)), gasPrice)

	fmt.Printf("Gas used: %d\nGas price: %s\nGas cost: %s\n", gasUsed, gasPrice, gasCost)
}
