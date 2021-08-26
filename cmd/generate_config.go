package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	defaultConfig = `api_key:
engine_id:
`
)

func newGenerateConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate_config_file",
		Short: "generate a default configuration file to $HOME/.config/lgotm/config",
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateConfig()
		},
	}
}

func generateConfig() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	path := filepath.Join(home, ".config", "lgotm")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm) // check mode
		if err != nil {
			return err
		}
	}

	configPath := filepath.Join(path, "config")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		f, err := os.Create(configPath)
		if err != nil {
			return err
		}
		defer f.Close()

		f.WriteString(defaultConfig)
	}
	return nil
}
