package cmd

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/skanehira/clipboard-image"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"

	"github.com/johnmanjiro13/lgotm/infra"
)

type QueryConfig struct {
	APIKey   string `mapstructure:"API_KEY"`
	EngineID string `mapstructure:"ENGINE_ID"`
}

func newQueryCmd() *cobra.Command {
	var cfgFile string
	var cfg QueryConfig

	queryCmd := &cobra.Command{
		Use:   "query",
		Short: "search images by query and generates a image which includes lgtm text",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return query(context.Background(), args, cfgFile, &cfg)
		},
	}

	queryCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/lgotm/config)")

	return queryCmd
}

type customSearchRepository interface {
	LGTM(context.Context, string) (io.Reader, error)
}

func query(ctx context.Context, args []string, cfgFile string, cfg *QueryConfig) error {
	initConfig(cfgFile, cfg)

	svc, err := customsearch.NewService(context.Background(), option.WithAPIKey(cfg.APIKey))
	if err != nil {
		return err
	}
	customSearchRepo := infra.NewCustomSearchRepository(svc, cfg.EngineID)
	c := queryCommand{customSearchRepo: customSearchRepo}
	query := strings.Join(args[:], " ")
	return c.exec(ctx, query)
}

type queryCommand struct {
	customSearchRepo customSearchRepository
}

func (c *queryCommand) exec(ctx context.Context, query string) error {
	img, err := c.customSearchRepo.LGTM(ctx, query)
	if err != nil {
		return err
	}

	if err := clipboard.CopyToClipboard(img); err != nil {
		return err
	}
	return nil
}

func initConfig(cfgFile string, cfg *QueryConfig) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(filepath.Join(home, ".config/lgotm"))
	}

	viper.AutomaticEnv()

	viper.BindEnv("api_key", "API_KEY")
	viper.BindEnv("engine_id", "ENGINE_ID")

	viper.ReadInConfig()

	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("failed to marshal config file: %w", err))
		os.Exit(1)
	}
}
