package main

import (
	"os"

	"github.com/fenghaojiang/ethereum-etl/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
