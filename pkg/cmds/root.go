package cmds

import (
	"flag"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:               "hook",
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	rootCmd.AddCommand(NewCmdVersion())
	rootCmd.AddCommand(NewCmdRun())

	return rootCmd
}
