package cmd

import (
	"github.com/fenghaojiang/ethereum-etl/common/log"
	"github.com/fenghaojiang/ethereum-etl/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func newEthereumETLCmd() *cobra.Command {

	etlCmd := &cobra.Command{
		Use:   "etl",
		Short: "etl jobs for ethereum like block chain",
	}

	var configPath string
	etlCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file (eg. job_cfg_example.json)")

	etlCmd.Run = func(cmd *cobra.Command, args []string) {
		log.Info("ethereum-etl start...")
		confs, err := config.LoadConfig(configPath)
		if err != nil {
			log.Error("load config failed", zap.Error(err))
			return
		}
		log.Info("load config succeed", zap.String("path", configPath), zap.Any("jobs config", confs))
		for i := range confs {
			// TODO Load Conf
			_ = i
		}

	}
	return etlCmd
}
