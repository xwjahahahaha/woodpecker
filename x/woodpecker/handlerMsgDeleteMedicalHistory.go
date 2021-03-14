package woodpecker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/xwj/woodpecker/x/woodpecker/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
)

// Handle a message to delete name
func handleMsgDeleteMedicalHistory(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteMedicalHistory) (*sdk.Result, error) {
	if !k.MedicalHistoryExists(ctx, msg.HashKey, msg.ID) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.ID)
	}
	if !msg.Creator.Equals(k.GetMedicalHistoryOwner(ctx, msg.HashKey, msg.ID)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteMedicalHistory(ctx, msg.HashKey, msg.ID)
	return &sdk.Result{}, nil
}
