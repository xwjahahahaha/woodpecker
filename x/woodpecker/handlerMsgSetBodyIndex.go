package woodpecker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/xwj/woodpecker/x/woodpecker/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
)

func handleMsgSetBodyIndex(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetBodyIndex) (*sdk.Result, error) {
	var bodyIndex = types.BodyIndex{
		Creator: msg.Creator,
    	Age: msg.Age,
    	Sex: msg.Sex,
    	Nation: msg.Nation,
    	Weight: msg.Weight,
    	Height: msg.Height,
    	WeightIndex: msg.WeightIndex,
    	ObesityWaistline: msg.ObesityWaistline,
    	Waistline: msg.Waistline,
    	MaxBloodPressure: msg.MaxBloodPressure,
    	MinBloodPressure: msg.MinBloodPressure,
    	GoodCholesterol: msg.GoodCholesterol,
    	BatCholesterol: msg.BatCholesterol,
    	TotalCholesterol: msg.TotalCholesterol,
    	Dyslipidemia: msg.Dyslipidemia,
    	Pvd: msg.Pvd,
    	SportActivities: msg.SportActivities,
    	Education: msg.Education,
    	Marry: msg.Marry,
    	Income: msg.Income,
    	SourceCase: msg.SourceCase,
    	VisionBad: msg.VisionBad,
    	Drink: msg.Drink,
    	HighBloodPressure: msg.HighBloodPressure,
    	FamilialHighBloodPressure: msg.FamilialHighBloodPressure,
    	Diabetes: msg.Diabetes,
    	FamilialDiabetes: msg.FamilialDiabetes,
    	Hepatitis: msg.Hepatitis,
    	FamilialHepatitis: msg.FamilialHepatitis,
    	ChronicFatigue: msg.ChronicFatigue,
    	HashKey: msg.HashKey,
	}
	if k.BodyIndexExists(ctx, msg.HashKey) && !msg.Creator.Equals(k.GetBodyIndexOwner(ctx, msg.HashKey)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetBodyIndex(ctx, bodyIndex)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
