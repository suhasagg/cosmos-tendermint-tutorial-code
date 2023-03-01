package keeper

import (
	"github.com/cosmos/cosmos-sdk/types"
)

type ModuleKeeper struct {
	transientStore TransientStore
}

func (mk *ModuleKeeper) ExpensiveCalculation(ctx types.Context) int {
	var result int
	// Perform expensive calculation
	// ...
	// Store result in the transient store
	mk.transientStore.SetTransient(ctx, result)
	return result
}

func (mk *ModuleKeeper) AnotherTransaction(ctx types.Context) {
	// Retrieve result from the transient store
	result, exists := mk.transientStore.GetTransient(ctx)
	if !exists {
		// Recalculate result if it doesn't exist in the transient store
		result = mk.ExpensiveCalculation(ctx)
	}
	// Use result in the execution of this transaction
	// ...
}
