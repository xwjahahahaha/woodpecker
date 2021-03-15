package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteBodyIndex{}

type MsgDeleteBodyIndex struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  HashKey string `json:"hash_key" yaml:"hash_key"`
}

func NewMsgDeleteBodyIndex(creator sdk.AccAddress, hashKey string) MsgDeleteBodyIndex {
  return MsgDeleteBodyIndex{
		Creator: creator,
		HashKey: hashKey,
	}
}

func (msg MsgDeleteBodyIndex) Route() string {
  return RouterKey
}

func (msg MsgDeleteBodyIndex) Type() string {
  return "DeleteBodyIndex"
}

func (msg MsgDeleteBodyIndex) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteBodyIndex) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteBodyIndex) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.HashKey) == 0{
	  return sdkerrors.Wrap(ErrBodyIndexParameterCantEmpty, "BodyIndex's parameter Can't Empty")
  }
  return nil
}