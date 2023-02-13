package cmd

import (
	"github.com/fenghaojiang/ethereum-etl/common/log"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ethereum-etl",
		Short: "ethereum block chain etl tool by fenghaojiang",
	}
	rootCmd.AddCommand(newEthereumETLCmd())
	return rootCmd
}

func Execute() (err error) {
	defer log.Sync()
	return NewRootCmd().Execute()
}
