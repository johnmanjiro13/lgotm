package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	generateConfigCmd = &cobra.Command{
		Use:   "generate_config_file",
		Short: "generate a default configuration file to $HOME/.config/lgotm/config.yaml",
		RunE:  generateConfig,
	}
)

func init() {
	rootCmd.AddCommand(generateConfigCmd)
}

const (
	defaultConfig = `api_key:
engine_id:
`
)

func generateConfig(cmd *cobra.Command, args []string) error {
	c := &generateConfigCommand{}
	return c.exec()
}

type generateConfigCommand struct{}

func (c *generateConfigCommand) exec() error {
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

	configPath := filepath.Join(path, "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		f, err := os.Create(filepath.Join(path, "config.yaml"))
		if err != nil {
			return err
		}
		defer f.Close()

		f.WriteString(defaultConfig)
	}
	return nil
}
