package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type BodyIndex struct {
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