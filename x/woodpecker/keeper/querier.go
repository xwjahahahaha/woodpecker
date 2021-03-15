package keeper

import (
  // this line is used by starport scaffolding # 1
	"github.com/xwj/woodpecker/x/woodpecker/types"
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for woodpecker clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListBodyIndex:
			return listBodyIndex(ctx, k)
		case types.QueryGetBodyIndex:
			return getBodyIndex(ctx, path[1:], k)
		case types.QueryListMedicalHistory:
			return listMedicalHistory(ctx, path[1:], k)
		case types.QueryGetMedicalHistory:
			return getMedicalHistory(ctx, path[1:], k)
		case types.QueryAllListMedicalHistory:
			return listAllMedicalHistory(ctx, k)
		case types.QueryListAttribute:
			return listAttribute(ctx, k)
		case types.QueryGetAttribute:
			return getAttribute(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown woodpecker query endpoint")
		}
	}
}
