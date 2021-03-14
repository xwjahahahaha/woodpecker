package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteAttribute{}

type MsgDeleteAttribute struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  HashKey string `json:"hashKey" yaml:"hashKey"`
}

func NewMsgDeleteAttribute(creator sdk.AccAddress, hashKey string) MsgDeleteAttribute {
  return MsgDeleteAttribute{
		HashKey: hashKey,
		Creator: creator,
	}
}

func (msg MsgDeleteAttribute) Route() string {
  return RouterKey
}

func (msg MsgDeleteAttribute) Type() string {
  return "DeleteAttribute"
}

func (msg MsgDeleteAttribute) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteAttribute) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteAttribute) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.HashKey) == 0{
	  return sdkerrors.Wrap(ErrAttributeParameterCantEmpty, "Attribute's parameter Can't Empty")
  }
  return nil
}