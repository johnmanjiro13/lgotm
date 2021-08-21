package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/johnmanjiro13/lgotm/infra"
	"github.com/skanehira/clipboard-image"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

type Config struct {
	CustomSearch struct {
		APIKey   string `mapstructure:"api_key"`
		EngineID string `mapstructure:"engine_id"`
	} `mapstructure:"custom_search"`
}

var (
	// Used for flags.
	cfgFile string
	config  Config

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

		viper.SetConfigFile("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(filepath.Join(home, ".config/lgotm"))
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("failed to read config file: %w", err))
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("failed to marshal config file: %w", err))
		os.Exit(1)
	}

	viper.AutomaticEnv()
}

func Execute() error {
	return rootCmd.Execute()
}

func lgtm(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("some querires are required.")
	}
	svc, err := customsearch.NewService(context.Background(), option.WithAPIKey(config.CustomSearch.APIKey))
	if err != nil {
		return err
	}
	customSearchRepo := infra.NewCustomSearchRepository(svc, config.CustomSearch.EngineID)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	img, err := customSearchRepo.LGTM(ctx, strings.Join(args[:], " "))
	if err != nil {
		return err
	}

	if err := clipboard.CopyToClipboard(img); err != nil {
		return err
	}
	return nil
}
