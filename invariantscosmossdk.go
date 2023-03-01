import (
    "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/x/bank"
)

func bankInvariant(k bank.Keeper) (string, sdk.Invariant {
    return "Total supply", func(ctx sdk.Context) (string, bool) {
        totalSupply := k.GetSupply(ctx).GetTotal()
        sum := sdk.ZeroInt()
        k.IterateAllBalances(ctx, func(_ sdk.AccAddress, balance sdk.Coin) bool {
            sum = sum.Add(balance.Amount)
            return false
        })
        if !sum.Equal(totalSupply) {
            return sdk.FormatInvariant(types.ModuleName, "total supply",
                "sum of account balances (%s) does not equal total supply (%s)", sum, totalSupply), true
        }
        return "", false
    }
}

func init() {
    bank.RegisterInvariants(
        bankInvariant,
    )
}
