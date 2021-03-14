package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/xwj/woodpecker/x/woodpecker/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
)

// GetAttributeCount get the total number of attribute
func (k Keeper) GetAttributeCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.AttributeCountPrefix)
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

// SetAttributeCount set the total number of attribute
func (k Keeper) SetAttributeCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.AttributeCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// GetAttribute returns the attribute information
func (k Keeper) GetAttribute(ctx sdk.Context, hashKey string) (types.Attribute, error) {
	store := ctx.KVStore(k.storeKey)
	var attribute types.Attribute
	byteKey := []byte(types.AttributePrefix + hashKey)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &attribute)
	if err != nil {
		return attribute, err
	}
	return attribute, nil
}

// SetAttribute sets a attribute
func (k Keeper) SetAttribute(ctx sdk.Context, attribute types.Attribute, hashKey string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(attribute)
	key := []byte(types.AttributePrefix + hashKey)
	store.Set(key, bz)
}

// DeleteAttribute deletes a attribute
func (k Keeper) DeleteAttribute(ctx sdk.Context, hashKey string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.AttributePrefix + hashKey))
}

// Get creator of the item
func (k Keeper) GetAttributeOwner(ctx sdk.Context, hashKey string) sdk.AccAddress {
	attribute, err := k.GetAttribute(ctx, hashKey)
	if err != nil {
		return nil
	}
	return attribute.Creator
}

func (k Keeper) HasOwner(ctx sdk.Context, hashKey string) bool {
	att, _ := k.GetAttribute(ctx, hashKey)
	return !att.Creator.Empty()
}

// Check if the key exists in the store
func (k Keeper) AttributeExists(ctx sdk.Context, hashKey string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.AttributePrefix + hashKey))
}


//
// Functions used by querier
//

func listAttribute(ctx sdk.Context, k Keeper) ([]byte, error) {
	var attributeList []types.Attribute
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.AttributePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var attribute types.Attribute
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &attribute)
		attributeList = append(attributeList, attribute)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, attributeList)
	return res, nil
}

func getAttribute(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	attribute, err := k.GetAttribute(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, attribute)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

