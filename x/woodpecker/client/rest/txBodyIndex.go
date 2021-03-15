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



type setBodyIndexRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Age string `json:"age"`
	Sex string `json:"sex"`
	Nation string `json:"nation"`
	Weight string `json:"weight"`
	Height string `json:"height"`
	WeightIndex string `json:"weightIndex"`
	ObesityWaistline string `json:"obesityWaistline"`
	Waistline string `json:"waistline"`
	MaxBloodPressure string `json:"maxBloodPressure"`
	MinBloodPressure string `json:"minBloodPressure"`
	GoodCholesterol string `json:"goodCholesterol"`
	BatCholesterol string `json:"batCholesterol"`
	TotalCholesterol string `json:"totalCholesterol"`
	Dyslipidemia string `json:"Dyslipidemia"`
	Pvd string `json:"pvd"`
	SportActivities string `json:"sportActivities"`
	Education string `json:"education"`
	Marry string `json:"marry"`
	Income string `json:"income"`
	SourceCase string `json:"sourceCase"`
	VisionBad string `json:"visionBad"`
	Drink string `json:"drink"`
	HighBloodPressure string `json:"highBloodPressure"`
	FamilialHighBloodPressure string `json:"familialHighBloodPressure"`
	Diabetes string `json:"diabetes"`
	FamilialDiabetes string `json:"familialDiabetes"`
	Hepatitis string `json:"hepatitis"`
	FamilialHepatitis string `json:"familialHepatitis"`
	ChronicFatigue string `json:"chronicFatigue"`
	HashKey string `json:"hash_key" yaml:"hash_key"`
	ALF string `json:"alf" yaml:"alf"`
}

func setBodyIndexHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = AllowCors(w)
		var req setBodyIndexRequest
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

		
		parsedAge64, err := strconv.ParseInt(req.Age, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedAge := int32(parsedAge64)
			
		parsedSex64, err := strconv.ParseInt(req.Sex, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedSex := int32(parsedSex64)
			
		parsedNation := req.Nation
		
		parsedWeight := req.Weight
		
		parsedHeight := req.Height
		
		parsedWeightIndex := req.WeightIndex
		
		parsedObesityWaistline := req.ObesityWaistline
		
		parsedWaistline := req.Waistline
		
		parsedMaxBloodPressure := req.MaxBloodPressure
		
		parsedMinBloodPressure := req.MinBloodPressure
		
		parsedGoodCholesterol := req.GoodCholesterol
		
		parsedBatCholesterol := req.BatCholesterol
		
		parsedTotalCholesterol := req.TotalCholesterol
		
		parsedDyslipidemia := req.Dyslipidemia
		
		parsedPvd := req.Pvd
		
		parsedSportActivities := req.SportActivities
		
		parsedEducation := req.Education
		
		parsedMarry64, err := strconv.ParseInt(req.Marry, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedMarry := int32(parsedMarry64)
			
		parsedIncome := req.Income
		
		parsedSourceCase := req.SourceCase
		
		parsedVisionBad := req.VisionBad
		
		parsedDrink := req.Drink
		
		parsedHighBloodPressure := req.HighBloodPressure
		
		parsedFamilialHighBloodPressure := req.FamilialHighBloodPressure
		
		parsedDiabetes := req.Diabetes
		
		parsedFamilialDiabetes := req.FamilialDiabetes
		
		parsedHepatitis := req.Hepatitis
		
		parsedFamilialHepatitis := req.FamilialHepatitis
		
		parsedChronicFatigue := req.ChronicFatigue

		parseALF := req.ALF

		parseHashKey := req.HashKey

		msg := types.NewMsgSetBodyIndex(
			creator,
			parsedAge,
			parsedSex,
			parsedNation,
			parsedWeight,
			parsedHeight,
			parsedWeightIndex,
			parsedObesityWaistline,
			parsedWaistline,
			parsedMaxBloodPressure,
			parsedMinBloodPressure,
			parsedGoodCholesterol,
			parsedBatCholesterol,
			parsedTotalCholesterol,
			parsedDyslipidemia,
			parsedPvd,
			parsedSportActivities,
			parsedEducation,
			parsedMarry,
			parsedIncome,
			parsedSourceCase,
			parsedVisionBad,
			parsedDrink,
			parsedHighBloodPressure,
			parsedFamilialHighBloodPressure,
			parsedDiabetes,
			parsedFamilialDiabetes,
			parsedHepatitis,
			parsedFamilialHepatitis,
			parsedChronicFatigue,
			parseALF,
			parseHashKey,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteBodyIndexRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	HashKey string `json:"hash_key"`
}

func deleteBodyIndexHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = AllowCors(w)
		var req deleteBodyIndexRequest
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
		msg := types.NewMsgDeleteBodyIndex(creator, req.HashKey)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
