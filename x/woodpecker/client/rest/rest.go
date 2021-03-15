package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers woodpecker-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/woodpecker/bodyIndex", listBodyIndexHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/bodyIndex/{key}", getBodyIndexHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/bodyIndex", setBodyIndexHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/woodpecker/bodyIndex", deleteBodyIndexHandler(cliCtx)).Methods("DELETE")

		


		r.HandleFunc("/woodpecker/medicalHistory", createMedicalHistoryHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/woodpecker/medicalHistory", listAllMedicalHistoryHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/medicalHistory/{hashKey}", listMedicalHistoryHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/medicalHistory/{hashKey}/{id}", getMedicalHistoryHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/medicalHistory", setMedicalHistoryHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/woodpecker/medicalHistory", deleteMedicalHistoryHandler(cliCtx)).Methods("DELETE")

		
		r.HandleFunc("/woodpecker/attribute", listAttributeHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/attribute/{key}", getAttributeHandler(cliCtx, "woodpecker")).Methods("GET")
		r.HandleFunc("/woodpecker/attribute", setAttributeHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/woodpecker/attribute", deleteAttributeHandler(cliCtx)).Methods("DELETE")
}
