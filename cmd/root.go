package cmd

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "lgotm",
	}

	rootCmd.AddCommand(newGenerateConfigCmd())

	return rootCmd
}

type Config struct {
	APIKey   string `mapstructure:"API_KEY"`
	EngineID string `mapstructure:"ENGINE_ID"`
}

func Execute() error {
	return newRootCmd().Execute()
}
