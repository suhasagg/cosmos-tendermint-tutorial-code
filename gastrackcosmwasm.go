package main

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
	cosmwasm "github.com/cosmos/cosmos-sdk/x/wasm"
)

func main() {
	// create a new codec
	cdc := codec.New()

	// define the transaction hash
	txHash := "0x<transaction hash>"

	// get the transaction
	res, err := queryTx(cdc, txHash)
	if err != nil {
		fmt.Println("Error getting transaction:", err)
		return
	}

	// decode the transaction
	var tx cosmwasm.WasmTx
	err = cdc.UnmarshalJSON(res, &tx)
	if err != nil {
		fmt.Println("Error decoding transaction:", err)
		return
	}
	
	// get the smart contract message
	var msg cosmwasm.WasmExecute
	err = cdc.UnmarshalJSON(tx.Data, &msg)
	if err != nil {
		fmt.Println("Error decoding message:", err)
		return
	}

	// get the smart contract operation
	var op cosmwasm.SmartContractOp
	err = cdc.UnmarshalJSON(msg.Data, &op)
	if err != nil {
		fmt.Println("Error decoding operation:", err)
		return
	}

	// get the gas used by the operation
	gasUsed := op.GasUsed

	// get the gas price of the operation
	gasPrice := op.GasPrice

	// calculate the gas cost of the operation
	gasCost := sdk.NewInt(gasUsed).Mul(gasPrice)

	fmt.Printf("Gas used: %d\nGas price: %s\nGas cost: %s\n", gasUsed, gasPrice, gasCost)
}

// queryTx queries a transaction by its hash
func queryTx(cdc *codec.Codec, hash string) ([]byte, error) {
	return flags.BroadcastCLI.QueryWithData("custom/tx", []byte(hash))
}
