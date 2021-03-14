package woodpecker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/xwj/woodpecker/x/woodpecker/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
)

func handleMsgSetMedicalHistory(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetMedicalHistory) (*sdk.Result, error) {
	var medicalHistory = types.MedicalHistory{
		Creator: msg.Creator,
		ID:      msg.ID,
    	MedicalInstitutionName: msg.MedicalInstitutionName,
    	DiagnosisTime: msg.DiagnosisTime,
    	DiagnosisDepartment: msg.DiagnosisDepartment,
    	DiagnosisType: msg.DiagnosisType,
    	DiagnosisDoctor: msg.DiagnosisDoctor,
    	DiagnosisResult: msg.DiagnosisResult,
    	TreatmentOptions: msg.TreatmentOptions,
	}
	if !msg.Creator.Equals(k.GetMedicalHistoryOwner(ctx, msg.HashKey, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetMedicalHistory(ctx, medicalHistory, msg.HashKey)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
