package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/xwj/woodpecker/x/woodpecker/types"
)

// GetMedicalHistoryCount get the total number of medicalHistory
func (k Keeper) GetMedicalHistoryCount(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.MedicalHistoryCountPrefix)
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

// SetMedicalHistoryCount set the total number of medicalHistory
func (k Keeper) SetMedicalHistoryCount(ctx sdk.Context, count int64)  {
	store := ctx.KVStore(k.storeKey)
	byteKey := []byte(types.MedicalHistoryCountPrefix)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateMedicalHistory creates a medicalHistory
func (k Keeper) CreateMedicalHistory(ctx sdk.Context, msg types.MsgCreateMedicalHistory, hashKey string) {
	// Create the medicalHistory
	count := k.GetMedicalHistoryCount(ctx)
    var medicalHistory = types.MedicalHistory{
        Creator: msg.Creator,
        ID:      strconv.FormatInt(count, 10),
        MedicalInstitutionName: msg.MedicalInstitutionName,
        DiagnosisTime: msg.DiagnosisTime,
        DiagnosisDepartment: msg.DiagnosisDepartment,
        DiagnosisType: msg.DiagnosisType,
        DiagnosisDoctor: msg.DiagnosisDoctor,
        DiagnosisResult: msg.DiagnosisResult,
        TreatmentOptions: msg.TreatmentOptions,
    }

	store := ctx.KVStore(k.storeKey)
	// 储存是 prefix + hashKey + ID
	key := []byte(types.MedicalHistoryPrefix + hashKey + medicalHistory.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(medicalHistory)
	store.Set(key, value)

	// Update medicalHistory count
    k.SetMedicalHistoryCount(ctx, count+1)
}

// GetMedicalHistory returns the medicalHistory information
func (k Keeper) GetMedicalHistory(ctx sdk.Context, hashKey string, id string) (types.MedicalHistory, error) {
	store := ctx.KVStore(k.storeKey)
	var medicalHistory types.MedicalHistory
	byteKey := []byte(types.MedicalHistoryPrefix + hashKey + id)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &medicalHistory)
	if err != nil {
		return medicalHistory, err
	}
	return medicalHistory, nil
}

// SetMedicalHistory sets a medicalHistory
func (k Keeper) SetMedicalHistory(ctx sdk.Context, medicalHistory types.MedicalHistory, hashKey string) {
	medicalHistoryKey := hashKey + medicalHistory.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(medicalHistory)
	key := []byte(types.MedicalHistoryPrefix + medicalHistoryKey)
	store.Set(key, bz)
}

// DeleteMedicalHistory deletes a medicalHistory
func (k Keeper) DeleteMedicalHistory(ctx sdk.Context, hashKey, id string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.MedicalHistoryPrefix + hashKey + id))
}


// Get creator of the item
func (k Keeper) GetMedicalHistoryOwner(ctx sdk.Context, hashKey, id string) sdk.AccAddress {
	medicalHistory, err := k.GetMedicalHistory(ctx, hashKey, id)
	if err != nil {
		return nil
	}
	return medicalHistory.Creator
}


// Check if the key exists in the store
func (k Keeper) MedicalHistoryExists(ctx sdk.Context, hashKey, id string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.MedicalHistoryPrefix + hashKey + id))
}


//
// Functions used by querier
//

func ListMedicalHistory(ctx sdk.Context, k Keeper, prefix string) ([]byte, error) {
	var medicalHistoryList []types.MedicalHistory
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.MedicalHistoryPrefix + prefix))
	for ; iterator.Valid(); iterator.Next() {
		var medicalHistory types.MedicalHistory
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &medicalHistory)
		medicalHistoryList = append(medicalHistoryList, medicalHistory)
	}

	res := codec.MustMarshalJSONIndent(k.cdc, medicalHistoryList)
	return res, nil
}

// 遍历一个人的所有病历
func listMedicalHistory(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	hashKey := path[0]
	return ListMedicalHistory(ctx, k, hashKey)
}


// 遍历所有人的所有病历
func listAllMedicalHistory(ctx sdk.Context, k Keeper,) ([]byte, error) {
	return ListMedicalHistory(ctx, k, "")
}

func getMedicalHistory(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	hashKey := path[0]
	ID := path[1]
	medicalHistory, err := k.GetMedicalHistory(ctx, hashKey, ID)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, medicalHistory)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
