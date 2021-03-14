package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/xwj/woodpecker/x/woodpecker/types"
)

func GetCmdSetAttribute(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-attribute [name] [idNumber] [address] [job] [nation] [hashKey]",
		Short: "Set a new attribute",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0] )
			argsIdNumber := string(args[1] )
			argsAddress := string(args[2] )
			argsJob := string(args[3] )
			argsNation := string(args[4] )
			argsHashKey := string(args[5])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetAttribute(cliCtx.GetFromAddress(), string(argsHashKey), string(argsName), string(argsIdNumber), string(argsAddress), string(argsJob), string(argsNation))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}



func GetCmdDeleteAttribute(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-attribute [hashKey]",
		Short: "Delete a new attribute by hashKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteAttribute(cliCtx.GetFromAddress(), args[0])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
