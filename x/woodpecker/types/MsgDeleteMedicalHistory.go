package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDeleteMedicalHistory{}

type MsgDeleteMedicalHistory struct {
  ID      string         `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  HashKey string `json:"hashKey" yaml:"hashKey"`
}

func NewMsgDeleteMedicalHistory(id string, creator sdk.AccAddress, hashKey string) MsgDeleteMedicalHistory {
  return MsgDeleteMedicalHistory{
    	ID: id,
		Creator: creator,
		HashKey: hashKey,
	}
}

func (msg MsgDeleteMedicalHistory) Route() string {
  return RouterKey
}

func (msg MsgDeleteMedicalHistory) Type() string {
  return "DeleteMedicalHistory"
}

func (msg MsgDeleteMedicalHistory) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgDeleteMedicalHistory) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgDeleteMedicalHistory) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.HashKey) == 0{
	return sdkerrors.Wrap(ErrMedicalHistoryParameterCantEmpty, "MedicalHistory's parameter Can't Empty")
  }
  return nil
}