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
	googleopt "google.golang.org/api/option"

	"github.com/johnmanjiro13/lgotm/image"
	"github.com/johnmanjiro13/lgotm/infra"
)

type QueryConfig struct {
	APIKey   string `mapstructure:"API_KEY"`
	EngineID string `mapstructure:"ENGINE_ID"`
}

type queryOption struct {
	cfg    *QueryConfig
	width  uint
	height uint
}

func newQueryCmd() *cobra.Command {
	var cfgFile string
	var cfg QueryConfig
	var width, height uint

	queryCmd := &cobra.Command{
		Use:   "query",
		Short: "search images by query and generates a image which includes lgtm text",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			initConfig(cfgFile, &cfg)
			opt := &queryOption{
				cfg:    &cfg,
				width:  width,
				height: height,
			}
			return query(context.Background(), args, opt)
		},
	}

	queryCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/lgotm/config)")
	queryCmd.Flags().UintVar(&width, "width", 0, "width of output image. the aspect ratio is kept if either width or height is 0")
	queryCmd.Flags().UintVar(&height, "height", 0, "height of output image. the aspect ratio is kept if either width or height is 0")

	return queryCmd
}

type customSearchRepository interface {
	FindImage(context.Context, string) (io.Reader, error)
}

func query(ctx context.Context, args []string, opt *queryOption) error {
	svc, err := customsearch.NewService(context.Background(), googleopt.WithAPIKey(opt.cfg.APIKey))
	if err != nil {
		return err
	}
	customSearchRepo := infra.NewCustomSearchRepository(svc, opt.cfg.EngineID)
	c := queryCommand{customSearchRepo: customSearchRepo}
	query := strings.Join(args[:], " ")
	return c.exec(ctx, query, opt.width, opt.height)
}

type queryCommand struct {
	customSearchRepo customSearchRepository
}

func (c *queryCommand) exec(ctx context.Context, query string, width, height uint) error {
	img, err := c.customSearchRepo.FindImage(ctx, query)
	if err != nil {
		return err
	}

	d := image.NewDrawer()
	res, err := d.LGTM(img, width, height)
	if err != nil {
		return err
	}

	if err := clipboard.CopyToClipboard(res); err != nil {
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
