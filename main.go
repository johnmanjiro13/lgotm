package main

import (
	"fmt"
	"os"

	"github.com/johnmanjiro13/lgotm/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
