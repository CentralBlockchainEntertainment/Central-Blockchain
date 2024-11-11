package encryption

import (
    "encoding/json"
    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/keeper"
    "github.com/CentralBlockchainEntertainment/Central-Blockchain/x/encryption/types"
    
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/x/simulation"
    abci "github.com/tendermint/tendermint/abci/types"
)

// AppModuleBasic defines the basic application module used by the encryption module.
type AppModuleBasic struct{}

func (AppModuleBasic) Name() string { return types.ModuleName }

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
    types.RegisterCodec(cdc)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
    return cdc.MustMarshalJSON(types.DefaultGenesis())
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ sdk.TxEncodingConfig, bz json.RawMessage) error {
    var genesisState types.GenesisState
    if err := cdc.UnmarshalJSON(bz, &genesisState); err != nil {
        return err
    }
    return types.ValidateGenesis(genesisState)
}

// RegisterInterfaces registers the module interfaces for protobuf compatibility.
func (AppModuleBasic) RegisterInterfaces(registry types.InterfaceRegistry) {
    types.RegisterInterfaces(registry)
}

// AppModule implements the AppModule interface for the encryption module.
type AppModule struct {
    AppModuleBasic
    keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(k keeper.Keeper) AppModule {
    return AppModule{
        AppModuleBasic: AppModuleBasic{},
        keeper:         k,
    }
}

func (am AppModule) Name() string { return types.ModuleName }

func (am AppModule) Route() sdk.Route {
    return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

func (am AppModule) QuerierRoute() string { return types.QuerierRoute }

func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
    return nil // Define querier functions if needed
}

// InitGenesis initializes the module's state from the genesis data.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
    var genesisState types.GenesisState
    cdc.MustUnmarshalJSON(data, &genesisState)
    InitGenesis(ctx, am.keeper, genesisState)
    return []abci.ValidatorUpdate{}
}

// ExportGenesis exports the module's state to the genesis data.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
    gs := ExportGenesis(ctx, am.keeper)
    return cdc.MustMarshalJSON(gs)
}

func (AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {}
func (AppModule) EndBlock(ctx sdk.Context, req abci.RequestEndBlock) []abci.ValidatorUpdate {
    return []abci.ValidatorUpdate{}
}
