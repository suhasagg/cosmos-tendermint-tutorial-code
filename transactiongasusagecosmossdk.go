package main

import (
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func main() {
	// define the codec to use for encoding and decoding messages
	cdc := codec.New()

	// create a new client context
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	// define the transaction message
	txBytes := []byte("<transaction bytes>")

	// decode the transaction message
	var msg auth.StdTx
	err := cdc.UnmarshalJSON(txBytes, &msg)
	if err != nil {
		fmt.Println("Error decoding transaction message:", err)
		return
	}

	// check if the transaction message is valid
	res, err := msg.ValidateBasic()
	if err != nil {
		fmt.Println("Transaction message is invalid:", err)
		return
	}

	if !res.IsOK() {
		fmt.Println("Transaction message is invalid:", res.Log)
		return
	}

	// calculate the gas usage of the transaction
	gasUsed := new(big.Int).SetUint64(msg.Fee.Gas)
	gasPrice := new(big.Int).SetUint64(msg.Fee.Amount.AmountOf(sdk.DefaultBondDenom))
	gasCost := new(big.Int).Mul(gasUsed, gasPrice)

	fmt.Printf("Gas used: %d\nGas price: %s\nGas cost: %s\n", gasUsed, gasPrice, gasCost)
}

