package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/johnmanjiro13/lgotm/cmd"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
