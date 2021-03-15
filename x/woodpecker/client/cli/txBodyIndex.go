package cli

import (
	"bufio"
    "strconv"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/xwj/woodpecker/x/woodpecker/types"
)

func GetCmdSetBodyIndex(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-bodyIndex [age] [sex] [nation] [weight] [height] [weightIndex] [obesityWaistline] [waistline] [maxBloodPressure] [minBloodPressure] [goodCholesterol] [batCholesterol] [totalCholesterol] [Dyslipidemia] [pvd] [sportActivities] [education] [marry] [income] [sourceCase] [visionBad] [drink] [highBloodPressure] [familialHighBloodPressure] [diabetes] [familialDiabetes] [hepatitis] [familialHepatitis] [chronicFatigue] [hashKey]",
		Short: "Set a new bodyIndex",
		Args:  cobra.ExactArgs(30),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAge, _ := strconv.ParseInt(args[0], 10, 64)
			argsSex, _ := strconv.ParseInt(args[1], 10, 64)
			argsNation := string(args[2])
			argsWeight := string(args[3])
			argsHeight := string(args[4])
			argsWeightIndex := string(args[5])
			argsObesityWaistline := string(args[6])
			argsWaistline := string(args[7])
			argsMaxBloodPressure := string(args[8])
			argsMinBloodPressure := string(args[9])
			argsGoodCholesterol := string(args[10])
			argsBatCholesterol := string(args[11])
			argsTotalCholesterol := string(args[12])
			argsDyslipidemia := string(args[13])
			argsPvd := string(args[14])
			argsSportActivities := string(args[15])
			argsEducation := string(args[16])
			argsMarry, _ := strconv.ParseInt(args[17], 10, 64)
			argsIncome := string(args[18])
			argsSourceCase := string(args[19])
			argsVisionBad := string(args[20])
			argsDrink := string(args[21])
			argsHighBloodPressure := string(args[22])
			argsFamilialHighBloodPressure := string(args[23])
			argsDiabetes := string(args[24])
			argsFamilialDiabetes := string(args[25])
			argsHepatitis := string(args[26])
			argsFamilialHepatitis := string(args[27])
			argsChronicFatigue := string(args[28])
			argsHashKey := string(args[29])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetBodyIndex(cliCtx.GetFromAddress(), int32(argsAge), int32(argsSex), string(argsNation), string(argsWeight), string(argsHeight), string(argsWeightIndex), string(argsObesityWaistline), string(argsWaistline), string(argsMaxBloodPressure), string(argsMinBloodPressure), string(argsGoodCholesterol), string(argsBatCholesterol), string(argsTotalCholesterol), string(argsDyslipidemia), string(argsPvd), string(argsSportActivities), string(argsEducation), int32(argsMarry), string(argsIncome), string(argsSourceCase), string(argsVisionBad), string(argsDrink), string(argsHighBloodPressure), string(argsFamilialHighBloodPressure), string(argsDiabetes), string(argsFamilialDiabetes), string(argsHepatitis), string(argsFamilialHepatitis), string(argsChronicFatigue), string(argsHashKey))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteBodyIndex(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-bodyIndex [hashKey]",
		Short: "Delete a new bodyIndex by hashKey",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteBodyIndex(cliCtx.GetFromAddress(), args[0])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
