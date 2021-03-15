package woodpecker

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/xwj/woodpecker/x/woodpecker/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
)

// Handle a message to delete name
func handleMsgDeleteBodyIndex(ctx sdk.Context, k keeper.Keeper, msg types.MsgDeleteBodyIndex) (*sdk.Result, error) {
	if !k.BodyIndexExists(ctx, msg.HashKey) {
		// replace with ErrKeyNotFound for 0.39+
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.HashKey)
	}
	if !msg.Creator.Equals(k.GetBodyIndexOwner(ctx, msg.HashKey)) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	k.DeleteBodyIndex(ctx, msg.HashKey)
	return &sdk.Result{}, nil
}
