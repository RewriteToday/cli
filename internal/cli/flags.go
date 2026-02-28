package cli

import "github.com/spf13/cobra"

type RenderOptions struct {
	Format  string
	NoColor bool
}

type InteractiveOptions struct {
	Interactive bool
	NoColor     bool
}

type InteractiveRenderOptions struct {
	Format      string
	Interactive bool
	NoColor     bool
}

func ReadRenderOptions(cmd *cobra.Command) RenderOptions {
	return RenderOptions{
		Format:  ReadStringFlag(cmd, "output"),
		NoColor: ReadBoolFlag(cmd, "no-color"),
	}
}

func ReadInteractiveOptions(cmd *cobra.Command) InteractiveOptions {
	return InteractiveOptions{
		Interactive: ReadBoolFlag(cmd, "interactive"),
		NoColor:     ReadBoolFlag(cmd, "no-color"),
	}
}

func ReadInteractiveRenderOptions(cmd *cobra.Command) InteractiveRenderOptions {
	return InteractiveRenderOptions{
		Format:      ReadStringFlag(cmd, "output"),
		Interactive: ReadBoolFlag(cmd, "interactive"),
		NoColor:     ReadBoolFlag(cmd, "no-color"),
	}
}

func ReadBoolFlag(cmd *cobra.Command, name string) bool {
	value, _ := cmd.Flags().GetBool(name)
	return value
}

func ReadStringFlag(cmd *cobra.Command, name string) string {
	value, _ := cmd.Flags().GetString(name)
	return value
}

func ReadIntFlag(cmd *cobra.Command, name string) int {
	value, _ := cmd.Flags().GetInt(name)
	return value
}
