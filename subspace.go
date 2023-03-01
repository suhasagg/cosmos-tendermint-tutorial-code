package types

import (
	"github.com/cosmos/cosmos-sdk/store/types"
)

const (
	DefaultParamspace types.SubspaceKey = "myapp"
)

var (
	KeyMyParameter = []byte("myparameter")
)

func ParamKey() []byte {
	return KeyMyParameter
}


import (
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/store/types"
)

func MyFunction(ctx sdk.Context) {
	// Get the store
	store := ctx.KVStore(myapp.DefaultParamspace)

	// Get the value of a parameter
	value := store.Get(KeyMyParameter)

	// Set the value of a parameter
	store.Set(KeyMyParameter, []byte("new value"))
}
