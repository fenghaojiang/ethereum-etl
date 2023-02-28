package cmd

import (
	"github.com/fenghaojiang/ethereum-etl/common/log"
	"github.com/spf13/cobra"
)

func newEthereumETLCmd() *cobra.Command {

	etlCmd := &cobra.Command{
		Use:   "etl",
		Short: "etl jobs for ethereum like block chain",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("ethereum-etl start...")

		},
	}
	return etlCmd
}
