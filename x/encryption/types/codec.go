package types

import (
    "github.com/cosmos/cosmos-sdk/codec"
    "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterCodec registers concrete message types for MsgEncrypt and MsgDecrypt
func RegisterCodec(cdc *codec.LegacyAmino) {
    cdc.RegisterConcrete(MsgEncrypt{}, "encryption/MsgEncrypt", nil)
    cdc.RegisterConcrete(MsgDecrypt{}, "encryption/MsgDecrypt", nil)
}

// RegisterInterfaces registers the module interfaces with the Cosmos SDK registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
    registry.RegisterImplementations((*sdk.Msg)(nil),
        &MsgEncrypt{},
        &MsgDecrypt{},
    )
}

// ModuleCdc is the codec for the encryption module, used for serialization
var (
    ModuleCdc = codec.NewAminoCodec(codec.NewLegacyAmino())
)
