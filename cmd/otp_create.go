package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandotp "github.com/RewriteToday/cli/internal/commands/otp"
	"github.com/spf13/cobra"
)

var otpCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create and send an OTP",
	Long:  "Create an OTP attempt through the Rewrite `/otp` API.",
	Example: `  rewrite otp create --to +5511999999999
  rewrite otp create --to +5511999999999 --prefix Rewrite --expires-in 10`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		return commandotp.Create(commandotp.CreateOpts{
			RenderOptions:  cliutil.ReadRenderOptions(cmd),
			To:             cliutil.ReadStringFlag(cmd, "to"),
			Prefix:         cliutil.ReadStringFlag(cmd, "prefix"),
			ExpiresIn:      cliutil.ReadIntFlag(cmd, "expires-in"),
			IdempotencyKey: cliutil.ReadStringFlag(cmd, "idempotency-key"),
		})
	},
}

func init() {
	otpCreateCmd.Flags().String("to", "", "Destination phone number in E.164 format")
	otpCreateCmd.Flags().String("prefix", "", "Optional OTP label/prefix")
	otpCreateCmd.Flags().Int("expires-in", 0, "Optional OTP expiration in minutes")
	otpCreateCmd.Flags().String("idempotency-key", "", "Optional idempotency key")

	otpCmd.AddCommand(otpCreateCmd)
}
