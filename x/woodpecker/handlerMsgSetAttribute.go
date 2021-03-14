package woodpecker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/xwj/woodpecker/x/woodpecker/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
)

func handleMsgSetAttribute(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetAttribute) (*sdk.Result, error) {
	var attribute = types.Attribute{
		Creator: msg.Creator,
    	Name: msg.Name,
    	IdNumber: msg.IdNumber,
    	Address: msg.Address,
    	Job: msg.Job,
    	Nation: msg.Nation,
	}
	// 如果已经存在那么就判断是否为本人
	if k.HasOwner(ctx, msg.HashKey) && !msg.Creator.Equals(k.GetAttributeOwner(ctx, msg.HashKey)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}
	k.SetAttribute(ctx, attribute, msg.HashKey)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
