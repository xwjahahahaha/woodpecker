package woodpecker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/xwj/woodpecker/x/woodpecker/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
)

// Handle a message to delete name
func handleMsgDeleteAttribute(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteAttribute) (*sdk.Result, error) {
	if !k.AttributeExists(ctx, msg.HashKey) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.HashKey)
	}
	if !msg.Creator.Equals(k.GetAttributeOwner(ctx, msg.HashKey)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteAttribute(ctx, msg.HashKey)
	return &sdk.Result{}, nil
}
