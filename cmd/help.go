package cmd

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/RewriteToday/cli/internal/render"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var configureHelpOnce sync.Once

func configureHelp() {
	configureHelpOnce.Do(func() {
		rootCmd.InitDefaultHelpCmd()
		applyHelpConfig(rootCmd)
	})
}

func applyHelpConfig(cmd *cobra.Command) {
	cmd.SetHelpFunc(func(command *cobra.Command, _ []string) {
		renderCommandHelp(command)
	})
	cmd.SetUsageFunc(func(command *cobra.Command) error {
		renderCommandHelp(command)
		return nil
	})

	for _, child := range cmd.Commands() {
		applyHelpConfig(child)
	}
}

func renderCommandHelp(cmd *cobra.Command) {
	out := cmd.OutOrStdout()
	noColor := readHelpNoColor(cmd)

	printed := false

	if printHeader(out, cmd, noColor) {
		printed = true
	}
	if printDescription(out, cmd, noColor, printed) {
		printed = true
	}
	if printUsage(out, cmd, noColor, printed) {
		printed = true
	}
	if printExamples(out, cmd, noColor, printed) {
		printed = true
	}
	if printAliases(out, cmd, noColor, printed) {
		printed = true
	}
	if printCommands(out, cmd, noColor, printed) {
		printed = true
	}
	if printFlags(out, "Flags", cmd.NonInheritedFlags(), noColor, printed) {
		printed = true
	}
	if printFlags(out, "Global Flags", cmd.InheritedFlags(), noColor, printed) {
		printed = true
	}
	printMore(out, noColor, printed)
}

func readHelpNoColor(cmd *cobra.Command) bool {
	noColor, err := cmd.Flags().GetBool("no-color")
	if err != nil {
		return false
	}

	return noColor
}

func printHeader(out io.Writer, cmd *cobra.Command, noColor bool) bool {
	if cmd == rootCmd {
		title := render.PaintAll("Rewrite", noColor, render.Bold, render.VesperGold)
		fmt.Fprintf(
			out,
			"%s %s  %s\n",
			render.Paint(">", render.VesperSubtle, noColor),
			title,
			render.Paint("v"+cmd.Version, render.VesperSubtle, noColor),
		)
		return true
	}

	fmt.Fprintln(out, render.PaintAll(cmd.CommandPath(), noColor, render.Bold, render.VesperBrand))
	return true
}

func printDescription(out io.Writer, cmd *cobra.Command, noColor bool, spaced bool) bool {
	description := strings.TrimSpace(cmd.Long)
	if description == "" {
		description = strings.TrimSpace(cmd.Short)
	}
	if description == "" {
		return false
	}

	if spaced {
		fmt.Fprintln(out)
	}

	for _, line := range strings.Split(description, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		fmt.Fprintln(out, render.Paint(line, render.VesperText, noColor))
	}

	return true
}

func printUsage(out io.Writer, cmd *cobra.Command, noColor bool, spaced bool) bool {
	if spaced {
		fmt.Fprintln(out)
	}

	fmt.Fprintf(out, "%s\n", paintSection("Usage", noColor))
	fmt.Fprintf(out, "  %s\n", render.Paint(helpUsageLine(cmd), render.VesperText, noColor))
	return true
}

func printExamples(out io.Writer, cmd *cobra.Command, noColor bool, spaced bool) bool {
	example := strings.TrimSpace(cmd.Example)
	if example == "" {
		return false
	}

	if spaced {
		fmt.Fprintln(out)
	}

	fmt.Fprintf(out, "%s\n", paintSection("Examples", noColor))
	for _, line := range strings.Split(example, "\n") {
		fmt.Fprintf(out, "  %s\n", render.Paint(strings.TrimSpace(line), render.VesperText, noColor))
	}

	return true
}

func printAliases(out io.Writer, cmd *cobra.Command, noColor bool, spaced bool) bool {
	if len(cmd.Aliases) == 0 {
		return false
	}

	if spaced {
		fmt.Fprintln(out)
	}

	fmt.Fprintf(out, "%s\n", paintSection("Aliases", noColor))
	fmt.Fprintf(out, "  %s\n", render.Paint(strings.Join(cmd.Aliases, ", "), render.VesperMuted, noColor))
	return true
}

func printCommands(out io.Writer, cmd *cobra.Command, noColor bool, spaced bool) bool {
	commands := visibleCommands(cmd)
	if len(commands) == 0 {
		return false
	}

	if spaced {
		fmt.Fprintln(out)
	}

	fmt.Fprintf(out, "%s\n", paintSection("Commands", noColor))

	maxWidth := 0
	for _, command := range commands {
		if width := len(command.Name()); width > maxWidth {
			maxWidth = width
		}
	}

	for _, command := range commands {
		name := command.Name()
		padding := strings.Repeat(" ", maxWidth-len(name)+2)
		fmt.Fprintf(
			out,
			"  %s%s%s\n",
			render.PaintAll(name, noColor, render.Bold, render.VesperGold),
			padding,
			render.Paint(command.Short, render.VesperText, noColor),
		)
	}

	return true
}

func printFlags(out io.Writer, title string, flags *pflag.FlagSet, noColor bool, spaced bool) bool {
	items := visibleFlags(flags, noColor)
	if len(items) == 0 {
		return false
	}

	if spaced {
		fmt.Fprintln(out)
	}

	fmt.Fprintf(out, "%s\n", paintSection(title, noColor))

	maxWidth := 0
	for _, item := range items {
		if width := len(item.label); width > maxWidth {
			maxWidth = width
		}
	}

	for _, item := range items {
		padding := strings.Repeat(" ", maxWidth-len(item.label)+2)
		fmt.Fprintf(
			out,
			"  %s%s%s\n",
			render.Paint(item.label, render.VesperMuted, noColor),
			padding,
			item.description,
		)
	}

	return true
}

type helpItem struct {
	label       string
	description string
}

func visibleCommands(cmd *cobra.Command) []*cobra.Command {
	children := cmd.Commands()
	items := make([]*cobra.Command, 0, len(children))

	for _, child := range children {
		if child.Hidden || !child.IsAvailableCommand() {
			continue
		}

		items = append(items, child)
	}

	return items
}

func visibleFlags(flags *pflag.FlagSet, noColor bool) []helpItem {
	if flags == nil {
		return nil
	}

	items := make([]helpItem, 0, flags.NFlag())
	flags.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		items = append(items, helpItem{
			label:       formatFlagLabel(flag),
			description: formatFlagDescription(flag, noColor),
		})
	})

	return items
}

func formatFlagLabel(flag *pflag.Flag) string {
	var parts [2]string
	count := 0

	if flag.Shorthand != "" {
		parts[count] = "-" + flag.Shorthand
		count++
	}

	parts[count] = "--" + flag.Name
	count++

	label := strings.Join(parts[:count], ", ")
	if flag.Value.Type() == "bool" {
		return label
	}

	return label + " <" + flagPlaceholder(flag) + ">"
}

func flagPlaceholder(flag *pflag.Flag) string {
	switch flag.Value.Type() {
	case "string":
		return "value"
	case "int", "int8", "int16", "int32", "int64":
		return "number"
	default:
		return flag.Value.Type()
	}
}

func formatFlagDescription(flag *pflag.Flag, noColor bool) string {
	description := render.Paint(flag.Usage, render.VesperText, noColor)
	if !shouldShowFlagDefault(flag) {
		return description
	}

	defaultValue := render.Paint("(default: "+flag.DefValue+")", render.VesperSubtle, noColor)
	return description + " " + defaultValue
}

func shouldShowFlagDefault(flag *pflag.Flag) bool {
	if flag.Value.Type() == "bool" {
		return false
	}

	return flag.DefValue != ""
}

func paintSection(title string, noColor bool) string {
	return render.Paint(title, render.VesperTeal, noColor)
}

func printMore(out io.Writer, noColor bool, spaced bool) {
	if spaced {
		fmt.Fprintln(out)
	}

	fmt.Fprintf(out, "%s\n", paintSection("More", noColor))
	fmt.Fprintf(out, "  %s\n", render.Paint("Use \"rewrite [command] --help\" to see deeper details.", render.VesperMuted, noColor))
}

func helpUsageLine(cmd *cobra.Command) string {
	if cmd.HasAvailableSubCommands() && !cmd.Runnable() {
		line := cmd.CommandPath() + " [command]"
		if hasVisibleFlags(cmd.NonInheritedFlags()) || hasVisibleFlags(cmd.InheritedFlags()) {
			line += " [flags]"
		}

		return line
	}

	return cmd.UseLine()
}

func hasVisibleFlags(flags *pflag.FlagSet) bool {
	if flags == nil {
		return false
	}

	visible := false
	flags.VisitAll(func(flag *pflag.Flag) {
		if flag.Hidden {
			return
		}

		visible = true
	})

	return visible
}
