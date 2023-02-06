
// Begin Block code:

func BeginBlocker(ctx sdk.Context, keeper Keeper) {
// Get the current block height
blockHeight := ctx.BlockHeight()

// Calculate the inflation rate for this block
inflationRate := keeper.GetInflationRate(blockHeight)

// Mint new tokens based on the inflation rate
newCoins := sdk.NewCoins(sdk.NewCoin(keeper.GetDenom(), inflationRate))
keeper.AddCoins(ctx, keeper.GetModuleAddress(ModuleName), newCoins)
}

// Inflation Rate code:

func GetInflationRate(blockHeight int64) sdk.Dec {
// Define the target inflation rate
targetInflation := sdk.NewDec(7)

// Calculate the actual inflation rate based on the block height
actualInflation := targetInflation.MulInt64(blockHeight)

// Return the actual inflation rate
return actualInflation
}
