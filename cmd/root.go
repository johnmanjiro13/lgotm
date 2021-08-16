package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "lgotm <query>",
		Short: "Lgotm googles image by query and generates a image which includes lgtm text",
		RunE:  lgtm,
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.lgotm)")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(filepath.Join(home, ".config/lgotm"))
	}

	viper.AutomaticEnv()
}

func Execute() error {
	return rootCmd.Execute()
}

func lgtm(cmd *cobra.Command, args []string) error {
	fmt.Println("LGTM")
	return nil
}
