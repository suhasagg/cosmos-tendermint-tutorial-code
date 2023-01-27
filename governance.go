package governance

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

// Governance module constants
const (
	ModuleName = "governance"
	RouterKey  = ModuleName
	StoreKey   = ModuleName
)

// Governance module messages
var (
	MsgSubmitProposal = gov.MsgSubmitProposal
	MsgDeposit        = gov.MsgDeposit
	MsgVote           = gov.MsgVote
)

// Governance module type
type Governance struct {
	gov.Keeper
}

// NewGovernance creates a new Governance module
func NewGovernance(keeper gov.Keeper) Governance {
	return Governance{Keeper: keeper}
}

// InitGenesis initializes the governance module's state from a genesis file
func (g Governance) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState gov.GenesisState
	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	return g.Keeper.InitGenesis(ctx, genesisState)
}

// ExportGenesis exports the governance module's state to a genesis file
func (g Governance) ExportGenesis(ctx sdk.Context) json.RawMessage {
	genesisState := g.Keeper.ExportGenesis(ctx)
	return ModuleCdc.MustMarshalJSON(genesisState)
}

// RegisterInvariants registers the governance module's invariants
func (g Governance) RegisterInvariants(ir sdk.InvariantRegistry) {
	g.Keeper.RegisterInvariants(ir)
}

// Route returns the governance module's message routing key
func (g Governance) Route() string {
	return RouterKey
}

// NewHandler returns a new governance module handler
func (g Governance) NewHandler() sdk.Handler {
	return gov.NewHandler(g.Keeper)
}

// QuerierRoute returns the governance module's querier route
func (g Governance) QuerierRoute() string {
	return QuerierRoute
}

// NewQuerierHandler returns a new governance module querier
func (g Governance) NewQuerierHandler() sdk.Querier {
	return gov.NewQuerier(g.Keeper)
}
