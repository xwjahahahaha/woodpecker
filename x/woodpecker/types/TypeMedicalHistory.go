package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MedicalHistory struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    MedicalInstitutionName string `json:"medicalInstitutionName" yaml:"medicalInstitutionName"`
    DiagnosisTime string `json:"diagnosisTime" yaml:"diagnosisTime"`
    DiagnosisDepartment string `json:"diagnosisDepartment" yaml:"diagnosisDepartment"`
    DiagnosisType string `json:"diagnosisType" yaml:"diagnosisType"`
    DiagnosisDoctor string `json:"diagnosisDoctor" yaml:"diagnosisDoctor"`
    DiagnosisResult string `json:"diagnosisResult" yaml:"diagnosisResult"`
    TreatmentOptions string `json:"treatmentOptions" yaml:"treatmentOptions"`
}