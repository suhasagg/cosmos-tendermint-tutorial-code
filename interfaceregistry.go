package main

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/ibc"
)

type handler struct {
	keeper ibc.Keeper
}

func NewHandler(k ibc.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case ibc.RegisterInterfacesMsg:
			return handleRegisterInterfaces(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized ibc message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleRegisterInterfaces(ctx sdk.Context, k ibc.Keeper, msg ibc.RegisterInterfacesMsg) sdk.Result {
	for _, msg := range msg.Msgs {
		if err := k.RegisterInterface(ctx, msg.Name, msg.Version, msg.Methods); err != nil {
			return err.Result()
		}
	}
	return sdk.Result{}
}

func main() {
	k := ibc.NewKeeper(nil)
	handler := NewHandler(k)
	msg := ibc.RegisterInterfacesMsg{
		Msgs: []ibc.InterfaceRegistration{
			{
				Name:    "transfer",
				Version: "1.0.0",
				Methods: []string{"transfer", "get_balance"},
			},
			{
				Name:    "issue",
				Version: "1.0.0",
				Methods: []string{"issue", "burn"},
			},
		},
	}
	ctx := sdk.NewContext(nil, sdk.Header{}, false, nil)
	result := handler(ctx, msg)
	if result.IsOK() {
		fmt.Println("Interfaces registered successfully")
	} else {
		fmt.Println("Failed to register interfaces:", result.Log)
	}
}

