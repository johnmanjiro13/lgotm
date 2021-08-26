package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "lgotm",
	}

	rootCmd.AddCommand(newGenerateConfigCmd())
	rootCmd.AddCommand(newQueryCmd())

	return rootCmd
}

func Execute() error {
	return newRootCmd().Execute()
}
