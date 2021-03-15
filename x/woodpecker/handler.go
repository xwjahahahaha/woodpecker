package woodpecker

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
	"github.com/xwj/woodpecker/x/woodpecker/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    	// this line is used by starport scaffolding # 1
		case types.MsgSetBodyIndex:
			return handleMsgSetBodyIndex(ctx, k, msg)
		case types.MsgDeleteBodyIndex:
			return handleMsgDeleteBodyIndex(ctx, k, msg)
		case types.MsgCreateMedicalHistory:
			return handleMsgCreateMedicalHistory(ctx, k, msg)
		case types.MsgSetMedicalHistory:
			return handleMsgSetMedicalHistory(ctx, k, msg)
		case types.MsgDeleteMedicalHistory:
			return handleMsgDeleteMedicalHistory(ctx, k, msg)
		case types.MsgSetAttribute:
			return handleMsgSetAttribute(ctx, k, msg)
		case types.MsgDeleteAttribute:
			return handleMsgDeleteAttribute(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
