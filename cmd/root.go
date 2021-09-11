package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "lgotm",
	}

	rootCmd.AddCommand(newVersionCmd(os.Stdout))
	rootCmd.AddCommand(newGenerateConfigCmd())
	rootCmd.AddCommand(newQueryCmd())
	rootCmd.AddCommand(newFileCmd())

	return rootCmd
}

func Execute() error {
	return newRootCmd().Execute()
}
