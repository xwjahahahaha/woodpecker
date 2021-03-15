package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetBodyIndex{}

type MsgSetBodyIndex struct {
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Age int32 `json:"age" yaml:"age"`
  Sex int32 `json:"sex" yaml:"sex"`
  Nation string `json:"nation" yaml:"nation"`
  Weight string `json:"weight" yaml:"weight"`
  Height string `json:"height" yaml:"height"`
  WeightIndex string `json:"weightIndex" yaml:"weightIndex"`
  ObesityWaistline string `json:"obesityWaistline" yaml:"obesityWaistline"`
  Waistline string `json:"waistline" yaml:"waistline"`
  MaxBloodPressure string `json:"maxBloodPressure" yaml:"maxBloodPressure"`
  MinBloodPressure string `json:"minBloodPressure" yaml:"minBloodPressure"`
  GoodCholesterol string `json:"goodCholesterol" yaml:"goodCholesterol"`
  BatCholesterol string `json:"batCholesterol" yaml:"batCholesterol"`
  TotalCholesterol string `json:"totalCholesterol" yaml:"totalCholesterol"`
  Dyslipidemia string `json:"Dyslipidemia" yaml:"Dyslipidemia"`
  Pvd string `json:"pvd" yaml:"pvd"`
  SportActivities string `json:"sportActivities" yaml:"sportActivities"`
  Education string `json:"education" yaml:"education"`
  Marry int32 `json:"marry" yaml:"marry"`
  Income string `json:"income" yaml:"income"`
  SourceCase string `json:"sourceCase" yaml:"sourceCase"`
  VisionBad string `json:"visionBad" yaml:"visionBad"`
  Drink string `json:"drink" yaml:"drink"`
  HighBloodPressure string `json:"highBloodPressure" yaml:"highBloodPressure"`
  FamilialHighBloodPressure string `json:"familialHighBloodPressure" yaml:"familialHighBloodPressure"`
  Diabetes string `json:"diabetes" yaml:"diabetes"`
  FamilialDiabetes string `json:"familialDiabetes" yaml:"familialDiabetes"`
  Hepatitis string `json:"hepatitis" yaml:"hepatitis"`
  FamilialHepatitis string `json:"familialHepatitis" yaml:"familialHepatitis"`
  ChronicFatigue string `json:"chronicFatigue" yaml:"chronicFatigue"`
  HashKey string `json:"hash_key" yaml:"hash_key"`
}

func NewMsgSetBodyIndex(creator sdk.AccAddress, age int32, sex int32, nation string, weight string, height string, weightIndex string, obesityWaistline string, waistline string, maxBloodPressure string, minBloodPressure string, goodCholesterol string, batCholesterol string, totalCholesterol string, Dyslipidemia string, pvd string, sportActivities string, education string, marry int32, income string, sourceCase string, visionBad string, drink string, highBloodPressure string, familialHighBloodPressure string, diabetes string, familialDiabetes string, hepatitis string, familialHepatitis string, chronicFatigue string, hashKey string) MsgSetBodyIndex {
  return MsgSetBodyIndex{
    Creator: creator,
    Age: age,
    Sex: sex,
    Nation: nation,
    Weight: weight,
    Height: height,
    WeightIndex: weightIndex,
    ObesityWaistline: obesityWaistline,
    Waistline: waistline,
    MaxBloodPressure: maxBloodPressure,
    MinBloodPressure: minBloodPressure,
    GoodCholesterol: goodCholesterol,
    BatCholesterol: batCholesterol,
    TotalCholesterol: totalCholesterol,
    Dyslipidemia: Dyslipidemia,
    Pvd: pvd,
    SportActivities: sportActivities,
    Education: education,
    Marry: marry,
    Income: income,
    SourceCase: sourceCase,
    VisionBad: visionBad,
    Drink: drink,
    HighBloodPressure: highBloodPressure,
    FamilialHighBloodPressure: familialHighBloodPressure,
    Diabetes: diabetes,
    FamilialDiabetes: familialDiabetes,
    Hepatitis: hepatitis,
    FamilialHepatitis: familialHepatitis,
    ChronicFatigue: chronicFatigue,
    HashKey: hashKey,
	}
}

func (msg MsgSetBodyIndex) Route() string {
  return RouterKey
}

func (msg MsgSetBodyIndex) Type() string {
  return "SetBodyIndex"
}

func (msg MsgSetBodyIndex) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetBodyIndex) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetBodyIndex) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  if len(msg.HashKey) == 0{
    return sdkerrors.Wrap(ErrBodyIndexParameterCantEmpty, "BodyIndex's parameter Can't Empty")
  }
  return nil
}