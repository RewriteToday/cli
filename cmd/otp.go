package cmd

import "github.com/spf13/cobra"

var otpCmd = &cobra.Command{
	Use:   "otp",
	Short: "Create and verify OTP codes",
	Long:  "Work against Rewrite `/otp` endpoints to create verification codes and confirm them.",
	Example: `  rewrite otp create --to +5511999999999
  rewrite otp verify 1234567890 --to +5511999999999 --code 123456`,
}

func init() {
	rootCmd.AddCommand(otpCmd)
}
