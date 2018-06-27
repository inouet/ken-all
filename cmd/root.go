package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
)

var output io.Writer

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ken-all",
		Short: "",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.AddCommand(newAddressCmd())
	rootCmd.AddCommand(newOfficeCmd())
	return rootCmd
}

func Execute() (exit exitcode.ExitCode) {
	exit = exitcode.Normal
	if err := newRootCmd().Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
}

func isValidOutputType(outputType string) bool {
	if outputType == "json" || outputType == "csv" || outputType == "tsv" {
		return true
	}
	return false
}
