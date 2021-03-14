package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetAttribute{}

type MsgSetAttribute struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
  IdNumber string `json:"idNumber" yaml:"idNumber"`
  Address string `json:"address" yaml:"address"`
  Job string `json:"job" yaml:"job"`
  Nation string `json:"nation" yaml:"nation"`
  HashKey string `json:"hashKey" yaml:"hashKey"`
}

func NewMsgSetAttribute(creator sdk.AccAddress, hashKey string, name string, idNumber string, address string, job string, nation string) MsgSetAttribute {
  return MsgSetAttribute{
	Creator: creator,
    Name: name,
    IdNumber: idNumber,
    Address: address,
    Job: job,
    Nation: nation,
    HashKey: hashKey,
	}
}

func (msg MsgSetAttribute) Route() string {
  return RouterKey
}

func (msg MsgSetAttribute) Type() string {
  return "SetAttribute"
}

func (msg MsgSetAttribute) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetAttribute) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetAttribute) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.HashKey) == 0 || len(msg.Name) == 0 || len(msg.IdNumber) == 0 || len(msg.Address) == 0 || len(msg.Nation) == 0 || len(msg.Job) == 0{
    return sdkerrors.Wrap(ErrAttributeParameterCantEmpty, "Attribute's parameter Can't Empty")
  }
  return nil
}