package slashing

import (
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/x/slashing/types"

    sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
    storeKey  sdk.StoreKey // Unexposed key to access store from sdk.Context
    cdc       *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the slashing Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
    return Keeper{
        storeKey:  storeKey,
        cdc:       cdc,
    }
}

// Slash the validator for infraction
func (k Keeper) Slash(ctx sdk.Context, validator sdk.ValAddress, infractionHeight int64, infractionPower int64) {
    // Get the validator information
    validatorInfo, found := k.GetValidatorInfo(ctx, validator)
    if !found {
        return
    }

    // Calculate the slash amount
    slashAmount := k.CalculateSlashAmount(validatorInfo, infractionHeight, infractionPower)

    // Deduct the slash amount from the validator's bond
    k.DeductBond(ctx, validator, slashAmount)
}

// Calculate the slash amount for the validator
func (k Keeper) CalculateSlashAmount(validator types.ValidatorInfo, infractionHeight int64, infractionPower int64) sdk.Coin {
    // Get the validator signing info
    signingInfo := validator.SigningInfo

    // Calculate the slash fraction
    slashFraction := k.SlashFraction(signingInfo.Height, signingInfo.StartHeight, infractionHeight)

    // Calculate the slash amount
    slashAmount := sdk.NewDecFromInt(validator.Tokens).Mul(slashFraction).TruncateInt()

    return sdk.NewCoin(validator.BondDenom, slashAmount)
}

// Deduct bond from validator
func (k Keeper) DeductBond(ctx sdk.Context, validator sdk.ValAddress, slashAmount sdk.Coin) {
    // Get the validator information
    validatorInfo, found := k.GetValidatorInfo(ctx, validator)
    if !found {
        return
    }

    // Deduct the bond
    validatorInfo.Bond = validatorInfo.Bond.Sub(slashAmount)

    // Set the new validator information
    k.SetValidatorInfo(ctx, validatorInfo)
}
