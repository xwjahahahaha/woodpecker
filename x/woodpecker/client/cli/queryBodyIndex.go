package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
    "github.com/xwj/woodpecker/x/woodpecker/types"
)

func GetCmdListBodyIndex(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-bodyIndex",
		Short: "list all bodyIndex",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListBodyIndex, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list BodyIndex\n%s\n", err.Error())
				return nil
			}
			var out []types.BodyIndex
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetBodyIndex(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-bodyIndex [hashKey]",
		Short: "Query a bodyIndex by hashKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetBodyIndex, key), nil)
			if err != nil {
				fmt.Printf("could not resolve bodyIndex %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.BodyIndex
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
