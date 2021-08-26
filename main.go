package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/johnmanjiro13/lgotm/cmd"
)

func main() {
	os.Exit(run())
}

func run() int {
	rand.Seed(time.Now().UnixNano())

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
