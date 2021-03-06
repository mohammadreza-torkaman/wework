package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mohammadreza-torkaman/wework/x/recruiter/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitJob() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-job [title] [description] [tags (separate tags with \\\",\\\" if it's more than one)\" ] [post-deadline] [job-deadline] [max-price] [location] [jobType] ",
		Short: "Broadcast message submit-job",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argDescription := args[1]
			argTags := strings.Split(args[2], ",")
			argPostDeadline := args[3]
			argJobDeadline := args[4]
			argMaxPrice := args[5]
			argLocation := args[6]
			argJobType, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitJob(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argDescription,
				argTags,
				argPostDeadline,
				argJobDeadline,
				argMaxPrice,
				argLocation,
				argJobType,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
