package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

func newVersionCmd(out io.Writer) *cobra.Command {
	c := &cobra.Command{
		Use:   "version",
		Short: "show version",
		Run: func(cmd *cobra.Command, args []string) {
			showVersion(cmd)
		},
	}
	c.SetOut(out)
	return c
}

func showVersion(cmd *cobra.Command) {
	cmd.Println(fmt.Sprintf("lgotm version %s", version))
}
