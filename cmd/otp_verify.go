package cmd

import (
	cliutil "github.com/RewriteToday/cli/internal/cli"
	commandotp "github.com/RewriteToday/cli/internal/commands/otp"
	"github.com/spf13/cobra"
)

var otpVerifyCmd = &cobra.Command{
	Use:     "verify [id]",
	Short:   "Verify an OTP code",
	Long:    "Submit a code to the Rewrite `/otp/:id/verify` endpoint.",
	Args:    cobra.ExactArgs(1),
	Example: `  rewrite otp verify 1234567890 --to +5511999999999 --code 123456`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return commandotp.Verify(commandotp.VerifyOpts{
			RenderOptions: cliutil.ReadRenderOptions(cmd),
			ID:            args[0],
			To:            cliutil.ReadStringFlag(cmd, "to"),
			Code:          cliutil.ReadStringFlag(cmd, "code"),
		})
	},
}

func init() {
	otpVerifyCmd.Flags().String("to", "", "Destination phone number in E.164 format")
	otpVerifyCmd.Flags().String("code", "", "Numeric OTP code")

	otpCmd.AddCommand(otpVerifyCmd)
}
