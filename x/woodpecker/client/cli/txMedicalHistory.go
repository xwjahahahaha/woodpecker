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

func GetCmdCreateMedicalHistory(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-medicalHistory [medicalInstitutionName] [diagnosisTime] [diagnosisDepartment] [diagnosisType] [diagnosisDoctor] [diagnosisResult] [treatmentOptions] [hashKey]",
		Short: "Creates a new medicalHistory",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsMedicalInstitutionName := string(args[0] )
			argsDiagnosisTime := string(args[1] )
			argsDiagnosisDepartment := string(args[2] )
			argsDiagnosisType := string(args[3] )
			argsDiagnosisDoctor := string(args[4] )
			argsDiagnosisResult := string(args[5] )
			argsTreatmentOptions := string(args[6] )
			argsHashKey := string(args[7])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateMedicalHistory(cliCtx.GetFromAddress(), string(argsMedicalInstitutionName), string(argsDiagnosisTime), string(argsDiagnosisDepartment), string(argsDiagnosisType), string(argsDiagnosisDoctor), string(argsDiagnosisResult), string(argsTreatmentOptions), string(argsHashKey))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetMedicalHistory(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-medicalHistory [id]  [medicalInstitutionName] [diagnosisTime] [diagnosisDepartment] [diagnosisType] [diagnosisDoctor] [diagnosisResult] [treatmentOptions] [hashKey]",
		Short: "Set a new medicalHistory",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsMedicalInstitutionName := string(args[1])
			argsDiagnosisTime := string(args[2])
			argsDiagnosisDepartment := string(args[3])
			argsDiagnosisType := string(args[4])
			argsDiagnosisDoctor := string(args[5])
			argsDiagnosisResult := string(args[6])
			argsTreatmentOptions := string(args[7])
			argsHashKey := string(args[8])
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetMedicalHistory(cliCtx.GetFromAddress(), id, string(argsMedicalInstitutionName), string(argsDiagnosisTime), string(argsDiagnosisDepartment), string(argsDiagnosisType), string(argsDiagnosisDoctor), string(argsDiagnosisResult), string(argsTreatmentOptions), string(argsHashKey))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteMedicalHistory(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-medicalHistory [id] [hashKey]",
		Short: "Delete a new medicalHistory by ID And HashKey",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteMedicalHistory(args[0], cliCtx.GetFromAddress(), args[1])
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
