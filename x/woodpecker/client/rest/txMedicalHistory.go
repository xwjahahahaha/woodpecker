package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xwj/woodpecker/x/woodpecker/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createMedicalHistoryRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	MedicalInstitutionName string `json:"medicalInstitutionName"`
	DiagnosisTime string `json:"diagnosisTime"`
	DiagnosisDepartment string `json:"diagnosisDepartment"`
	DiagnosisType string `json:"diagnosisType"`
	DiagnosisDoctor string `json:"diagnosisDoctor"`
	DiagnosisResult string `json:"diagnosisResult"`
	TreatmentOptions string `json:"treatmentOptions"`
	HashKey string `json:"hash_key"`
}

func createMedicalHistoryHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createMedicalHistoryRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedMedicalInstitutionName := req.MedicalInstitutionName
		
		parsedDiagnosisTime := req.DiagnosisTime
		
		parsedDiagnosisDepartment := req.DiagnosisDepartment
		
		parsedDiagnosisType := req.DiagnosisType
		
		parsedDiagnosisDoctor := req.DiagnosisDoctor
		
		parsedDiagnosisResult := req.DiagnosisResult
		
		parsedTreatmentOptions := req.TreatmentOptions

		parsedHashKey := req.HashKey

		msg := types.NewMsgCreateMedicalHistory(
			creator,
			parsedMedicalInstitutionName,
			parsedDiagnosisTime,
			parsedDiagnosisDepartment,
			parsedDiagnosisType,
			parsedDiagnosisDoctor,
			parsedDiagnosisResult,
			parsedTreatmentOptions,
			parsedHashKey,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setMedicalHistoryRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Creator string `json:"creator"`
	MedicalInstitutionName string `json:"medicalInstitutionName"`
	DiagnosisTime string `json:"diagnosisTime"`
	DiagnosisDepartment string `json:"diagnosisDepartment"`
	DiagnosisType string `json:"diagnosisType"`
	DiagnosisDoctor string `json:"diagnosisDoctor"`
	DiagnosisResult string `json:"diagnosisResult"`
	TreatmentOptions string `json:"treatmentOptions"`
	HashKey string `json:"hash_key"`
}

func setMedicalHistoryHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setMedicalHistoryRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedMedicalInstitutionName := req.MedicalInstitutionName
		
		parsedDiagnosisTime := req.DiagnosisTime
		
		parsedDiagnosisDepartment := req.DiagnosisDepartment
		
		parsedDiagnosisType := req.DiagnosisType
		
		parsedDiagnosisDoctor := req.DiagnosisDoctor
		
		parsedDiagnosisResult := req.DiagnosisResult
		
		parsedTreatmentOptions := req.TreatmentOptions

		parsedHashKey := req.HashKey

		msg := types.NewMsgSetMedicalHistory(
			creator,
			req.ID,
			parsedMedicalInstitutionName,
			parsedDiagnosisTime,
			parsedDiagnosisDepartment,
			parsedDiagnosisType,
			parsedDiagnosisDoctor,
			parsedDiagnosisResult,
			parsedTreatmentOptions,
			parsedHashKey,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteMedicalHistoryRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID 		string `json:"id"`
	HashKey string `json:"hash_key"`
}

func deleteMedicalHistoryHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteMedicalHistoryRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgDeleteMedicalHistory(req.ID, creator, req.HashKey)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
