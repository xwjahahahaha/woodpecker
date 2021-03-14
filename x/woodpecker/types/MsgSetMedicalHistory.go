package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetMedicalHistory{}

type MsgSetMedicalHistory struct {
  ID      string      `json:"id" yaml:"id"`
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  MedicalInstitutionName string `json:"medicalInstitutionName" yaml:"medicalInstitutionName"`
  DiagnosisTime string `json:"diagnosisTime" yaml:"diagnosisTime"`
  DiagnosisDepartment string `json:"diagnosisDepartment" yaml:"diagnosisDepartment"`
  DiagnosisType string `json:"diagnosisType" yaml:"diagnosisType"`
  DiagnosisDoctor string `json:"diagnosisDoctor" yaml:"diagnosisDoctor"`
  DiagnosisResult string `json:"diagnosisResult" yaml:"diagnosisResult"`
  TreatmentOptions string `json:"treatmentOptions" yaml:"treatmentOptions"`
  HashKey string `json:"hashKey" yaml:"hashKey"`
}

func NewMsgSetMedicalHistory(creator sdk.AccAddress, id string, medicalInstitutionName string, diagnosisTime string, diagnosisDepartment string, diagnosisType string, diagnosisDoctor string, diagnosisResult string, treatmentOptions string, hashKey string) MsgSetMedicalHistory {
  return MsgSetMedicalHistory{
    ID: id,
    Creator: creator,
    MedicalInstitutionName: medicalInstitutionName,
    DiagnosisTime: diagnosisTime,
    DiagnosisDepartment: diagnosisDepartment,
    DiagnosisType: diagnosisType,
    DiagnosisDoctor: diagnosisDoctor,
    DiagnosisResult: diagnosisResult,
    TreatmentOptions: treatmentOptions,
    HashKey: hashKey,
	}
}

func (msg MsgSetMedicalHistory) Route() string {
  return RouterKey
}

func (msg MsgSetMedicalHistory) Type() string {
  return "SetMedicalHistory"
}

func (msg MsgSetMedicalHistory) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetMedicalHistory) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetMedicalHistory) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.HashKey) == 0{
    return sdkerrors.Wrap(ErrMedicalHistoryParameterCantEmpty, "MedicalHistory's parameter Can't Empty")
  }
  return nil
}