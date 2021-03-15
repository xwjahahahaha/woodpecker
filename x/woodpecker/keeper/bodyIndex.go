package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/xwj/woodpecker/x/woodpecker/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// GetBodyIndexCount get the total number of bodyIndex
func (k Keeper) GetBodyIndexCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.BodyIndexCountPrefix)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetBodyIndexCount set the total number of bodyIndex
func (k Keeper) SetBodyIndexCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.BodyIndexCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// GetBodyIndex returns the bodyIndex information
func (k Keeper) GetBodyIndex(ctx sdk.Context, hashKey string) (types.BodyIndex, error) {
	store := ctx.KVStore(k.storeKey)
	var bodyIndex types.BodyIndex
	byteKey := []byte(types.BodyIndexPrefix + hashKey)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &bodyIndex)
	if err != nil {
		return bodyIndex, err
	}
	return bodyIndex, nil
}

// SetBodyIndex sets a bodyIndex
func (k Keeper) SetBodyIndex(ctx sdk.Context, bodyIndex types.BodyIndex) {
	bodyIndexKey := bodyIndex.HashKey
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(bodyIndex)
	key := []byte(types.BodyIndexPrefix + bodyIndexKey)
	store.Set(key, bz)
}

// DeleteBodyIndex deletes a bodyIndex
func (k Keeper) DeleteBodyIndex(ctx sdk.Context, hashKey string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.BodyIndexPrefix + hashKey))
}

// Get creator of the item
func (k Keeper) GetBodyIndexOwner(ctx sdk.Context, hashKey string) sdk.AccAddress {
	bodyIndex, err := k.GetBodyIndex(ctx, hashKey)
	if err != nil {
		return nil
	}
	return bodyIndex.Creator
}


// Check if the key exists in the store
func (k Keeper) BodyIndexExists(ctx sdk.Context, hashKey string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.BodyIndexPrefix + hashKey))
}



//
// Functions used by querier
//

func listBodyIndex(ctx sdk.Context, k Keeper) ([]byte, error) {
	var bodyIndexList []types.BodyIndex
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.BodyIndexPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var bodyIndex types.BodyIndex
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &bodyIndex)
		bodyIndexList = append(bodyIndexList, bodyIndex)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, bodyIndexList)
	return res, nil
}

func getBodyIndex(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	bodyIndex, err := k.GetBodyIndex(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, bodyIndex)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

