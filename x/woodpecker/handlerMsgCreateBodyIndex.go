package woodpecker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/xwj/woodpecker/x/woodpecker/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
)

func handleMsgCreateBodyIndex(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateBodyIndex) (*sdk.Result, error) {
	k.CreateBodyIndex(ctx, msg)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
