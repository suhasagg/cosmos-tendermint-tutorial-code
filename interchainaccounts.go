package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

func main() {
	// Connect to the host zone
	ctx := context.NewCLIContext().WithNodeURI("<host_zone_uri>")

	// Get the account address and sequence number
	address, sequence, err := ctx.GetFromAddress()
	if err != nil {
		fmt.Println("Error getting account information:", err)
		return
	}

	// Define the interchain account details
	account := auth.BaseAccount{
		Address: address,
		Coins: []auth.Coin{
			{
				Denom:  "stake",
				Amount: 100,
			},
		},
		Sequence: sequence,
	}

	// Send the interchain account creation request to the controller zone
	res, err := bank.SendRequest(ctx, "<controller_zone_uri>", &account)
	if err != nil {
		fmt.Println("Error creating interchain account:", err)
		return
	}

	// Check the result of the interchain account creation
	if res.Code != 0 {
		fmt.Println("Interchain account creation failed:", res.Log)
		return
	}

	fmt.Println("Interchain account created successfully")
}
