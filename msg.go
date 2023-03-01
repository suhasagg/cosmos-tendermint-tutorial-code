package mymodule

import (
	"github.com/cosmos/cosmos-sdk/types"
)

// Define the module name
const ModuleName = "mymodule"

// Define a custom message type
type MyMsg struct {
	Sender types.AccAddress
	Recipient types.AccAddress
	Amount types.Coins
}

// Implement the sdk.Msg interface for the custom message type
func (msg MyMsg) Route() string {
	return ModuleName
}

func (msg MyMsg) Type() string {
	return "mymsg"
}

func (msg MyMsg) ValidateBasic() error {
	if msg.Sender.Empty() {
		return types.ErrInvalidAddress
	}
	if msg.Recipient.Empty() {
		return types.ErrInvalidAddress
	}
	if !msg.Amount.IsValid() {
		return types.ErrInvalidCoins
	}
	return nil
}
