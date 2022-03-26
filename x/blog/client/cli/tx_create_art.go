package cli

import (
	"strconv"

	"github.com/cosmonaut/blog/x/blog/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateArt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-art [title] [body] [price] [old-art] [author-art] [author-sender] [published] [image]",
		Short: "Broadcast message createArt",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argBody := args[1]
			argPrice := args[2]
			argOldArt := args[3]
			argAuthorArt := args[4]
			argAuthorSender := args[5]
			argPublished := args[6]
			argImage := args[7]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateArt(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argBody,
				argPrice,
				argOldArt,
				argAuthorArt,
				argAuthorSender,
				argPublished,
				argImage,
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
