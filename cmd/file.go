package cmd

import (
	"io"
	"os"

	"github.com/johnmanjiro13/lgotm/image"
	"github.com/skanehira/clipboard-image"
	"github.com/spf13/cobra"
)

func newFileCmd() *cobra.Command {
	var width, height uint

	fileCmd := &cobra.Command{
		Use:   "file",
		Short: "generates LGTM image from given file",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := &fileCommand{}
			img, err := c.lgtm(args[0], width, height)
			if err != nil {
				return err
			}
			if err := clipboard.CopyToClipboard(img); err != nil {
				return err
			}
			return nil
		},
	}

	fileCmd.Flags().UintVar(&width, "width", 0, "width of output image. the aspect ratio is kept if either width or height is 0")
	fileCmd.Flags().UintVar(&height, "height", 0, "height of output image. the aspect ratio is kept if either width or height is 0")

	return fileCmd
}

type fileCommand struct{}

func (c *fileCommand) lgtm(path string, width, height uint) (io.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := image.NewDrawer()
	res, err := d.LGTM(file, width, height)
	if err != nil {
		return nil, err
	}
	return res, nil
}
