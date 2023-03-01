package mymodule

import (
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/cosmos/cosmos-sdk/x/auth"
    "github.com/cosmos/cosmos-sdk/x/bank"
    "github.com/cosmos/cosmos-sdk/x/params"
)

var (
    _ module.AppModule      = AppModule{}
    _ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct{}

func (a AppModuleBasic) Name() string {
    return "mymodule"
}

func (a AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
    // Register your module's types here
}

func (a AppModuleBasic) DefaultGenesis() json.RawMessage {
    // Define your module's default genesis state here
}

func (a AppModuleBasic) ValidateGenesis(data json.RawMessage) error {
    // Validate your module's genesis state here
}

func (a AppModuleBasic) RegisterRESTRoutes(ctx context.Context, rtr *mux.Router) {
    // Register your module's REST endpoints here
}

func (a AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
    // Define your module's transaction command here
}

func (a AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
    // Define your module's query command here
}

type AppModule struct {
    AppModuleBasic
    keeper         Keeper
    accountKeeper  auth.AccountKeeper
    bankKeeper     bank.Keeper
    paramsKeeper   params.Keeper
}

func NewAppModule(keeper Keeper, ak auth.AccountKeeper, bk bank.Keeper, pk params.Keeper) AppModule {
    return AppModule{
        AppModuleBasic: AppModuleBasic{},
        keeper:         keeper,
        accountKeeper:  ak,
        bankKeeper:     bk,
        paramsKeeper:   pk,
    }
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
    // Register your module's invariants here
}

func (am AppModule) Route() sdk.Route {
    // Define your module's routes here
}

func (am AppModule) QuerierRoute() string {
    // Define your module's querier route here
}

func (am AppModule) LegacyQuerierHandler(cdc *codec.LegacyAmino) sdk.Querier {
    // Define your module's legacy querier handler here
}

func (am AppModule) RegisterServices(Configurator) {
    // Register your module's services here
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
    // Initialize your module's genesis state here
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
    // Export your module's genesis state here
}

func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
    // Define your module's BeginBlock logic here
}

func (am AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
    // Define your module's EndBlock logic here
}
