package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers woodpecker-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/woodpecker/attribute", listAttributeHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/attribute/{key}", getAttributeHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/attribute", setAttributeHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/woodpecker/attribute", deleteAttributeHandler(cliCtx)).Methods("DELETE")
}
