package common

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xwj/woodpecker/x/woodpecker/keeper"
	"github.com/xwj/woodpecker/x/woodpecker/types"
)

func ListType(ctx sdk.Context, k keeper.Keeper, prefix string) ([]byte, error) {
	var attributeList []types.Attribute
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(prefix))
	for ; iterator.Valid(); iterator.Next() {
		var attribute types.Attribute
		k.Cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &attribute)
		attributeList = append(attributeList, attribute)
	}
	res := codec.MustMarshalJSONIndent(k.Cdc, attributeList)
	return res, nil
}