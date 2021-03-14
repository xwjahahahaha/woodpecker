package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
    "github.com/xwj/woodpecker/x/woodpecker/types"
)

func GetCmdListAllMedicalHistory(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-all-medicalHistory",
		Short: "list all medicalHistory",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryAllListMedicalHistory, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list MedicalHistory\n%s\n", err.Error())
				return nil
			}
			var out []types.MedicalHistory
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}


func GetCmdListMedicalHistory(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-medicalHistory [hashKey]",
		Short: "list all medicalHistory by one user's key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			hashKey := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryListMedicalHistory, hashKey), nil)
			if err != nil {
				fmt.Printf("could not resolve medicalHistory %s \n%s\n", hashKey, err.Error())

				return nil
			}

			var out []types.MedicalHistory
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}


func GetCmdGetMedicalHistory(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-medicalHistory [hashKey] [id]",
		Short: "Query a medicalHistory by key and id",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			hashKey := args[0]
			id := args[1]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s/%s", queryRoute, types.QueryGetMedicalHistory, hashKey, id), nil)
			if err != nil {
				fmt.Printf("could not resolve medicalHistory %s \n%s\n", hashKey + " : " + id, err.Error())

				return nil
			}

			var out types.MedicalHistory
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
