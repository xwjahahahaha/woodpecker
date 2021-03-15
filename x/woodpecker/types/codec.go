package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgSetBodyIndex{}, "woodpecker/SetBodyIndex", nil)
		cdc.RegisterConcrete(MsgDeleteBodyIndex{}, "woodpecker/DeleteBodyIndex", nil)
		cdc.RegisterConcrete(MsgCreateMedicalHistory{}, "woodpecker/CreateMedicalHistory", nil)
		cdc.RegisterConcrete(MsgSetMedicalHistory{}, "woodpecker/SetMedicalHistory", nil)
		cdc.RegisterConcrete(MsgDeleteMedicalHistory{}, "woodpecker/DeleteMedicalHistory", nil)
		cdc.RegisterConcrete(MsgSetAttribute{}, "woodpecker/SetAttribute", nil)
		cdc.RegisterConcrete(MsgDeleteAttribute{}, "woodpecker/DeleteAttribute", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
