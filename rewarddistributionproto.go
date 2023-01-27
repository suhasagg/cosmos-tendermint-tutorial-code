package distribution

import (
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/x/distribution/types"

    sdk "github.com/cosmos/cosmos-sdk/types"
    abci "github.com/tendermint/tendermint/abci/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
    storeKey  sdk.StoreKey // Unexposed key to access store from sdk.Context
    cdc       *codec.Codec // The wire codec for binary encoding/decoding.
    paramspace types.ParamSubspace
}

// NewKeeper creates new instances of the distribution Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, paramspace types.ParamSubspace) Keeper {
    return Keeper{
        storeKey:  storeKey,
        cdc:       cdc,
        paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
    }
}

// DistributeFromFeePool distributes funds from the fee pool to validators and delegators
func (k Keeper) DistributeFromFeePool(ctx sdk.Context) {
    // Get the current fee pool
    feePool := k.GetFeePool(ctx)

    // Get the current distribution parameters
    params := k.GetParams(ctx)

    // Calculate the total distribution
    totalDistribution := feePool.CommunityPool.Add(feePool.InflationPool)

    // Distribute funds to validators
    k.DistributeToValidators(ctx, totalDistribution, params)

    // Distribute funds to delegators
    k.DistributeToDelegators(ctx, totalDistribution, params)

    // Reset the fee pool
    k.ResetFeePool(ctx)
}

// DistributeToValidators distributes funds to validators based on their outstanding shares
func (k Keeper) DistributeToValidators(ctx sdk.Context, distribution sdk.Coin, params types.Params) {
    // Get the current validator set
    validatorSet := k.GetValidatorSet(ctx)

    // Calculate the distribution per share
    distributionPerShare := distribution.ToDec().Quo(validatorSet.TotalShares.ToDec())

    // Iterate through the validators and distribute funds
    for _, validator := range validatorSet.Validators {
        // Calculate the distribution for this validator
        validatorDistribution := validator.Shares.MulDec(distributionPerShare)

        // Send funds to the validator
        k.SendCoinsFromModuleToAccount(ctx, types.ModuleName, validator.OperatorAddress, validatorDistribution)
    }
}
