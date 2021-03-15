package rest

import (
	"fmt"
	"github.com/xwj/woodpecker/x/woodpecker/types"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

)

func listAllMedicalHistoryHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = AllowCors(w)
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", storeName, types.QueryAllListMedicalHistory), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func listMedicalHistoryHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = AllowCors(w)
		vars := mux.Vars(r)
		key := vars["hashKey"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", storeName, types.QueryListMedicalHistory, key), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}


func getMedicalHistoryHandler(cliCtx context.CLIContext, storeName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = AllowCors(w)
		vars := mux.Vars(r)
		hashKey := vars["hashKey"]
		ID := vars["id"]
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s/%s", storeName, types.QueryGetMedicalHistory, hashKey, ID), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
