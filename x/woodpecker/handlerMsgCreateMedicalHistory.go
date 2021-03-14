package woodpecker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/xwj/woodpecker/x/woodpecker/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
)

func handleMsgCreateMedicalHistory(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateMedicalHistory) (*sdk.Result, error) {
	k.CreateMedicalHistory(ctx, msg, msg.HashKey)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
