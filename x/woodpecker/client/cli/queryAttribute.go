package cli

import (
	"fmt"
	"github.com/xwj/woodpecker/x/woodpecker/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

)

func GetCmdListAttribute(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-attribute",
		Short: "list all attribute",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListAttribute, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Attribute\n%s\n", err.Error())
				return nil
			}
			var out []types.Attribute
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetAttribute(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-attribute [hashKey]",
		Short: "Query a attribute by hashKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			key := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetAttribute, key), nil)
			if err != nil {
				fmt.Printf("could not resolve attribute %s \n%s\n", key, err.Error())

				return nil
			}

			var out types.Attribute
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
