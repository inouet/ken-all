package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ken-all",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(newConvertCmd())
	return rootCmd
}

func Execute() (exit exitcode.ExitCode) {
	exit = exitcode.Normal
	if err := newRootCmd().Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
}
