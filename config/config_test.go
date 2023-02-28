package config

import (
	"testing"

	"github.com/fenghaojiang/ethereum-etl/common/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestOnLoadConfig(t *testing.T) {
	t.Run("test on load config", func(t *testing.T) {
		jobs, err := LoadConfig("../job_cfg_example.json")
		if err != nil {
			log.Error("failed to load config", zap.Error(err))
			t.FailNow()
		}
		assert.Equal(t, 1, len(jobs))

		assert.Equal(t, JobConf{
			JobID:      1,
			Network:    "ethereum",
			ETLJobType: "transactions",
			RPCEndpoints: []string{
				"http://127.0.0.1:8545",
				"wss://127.0.0.1:8545",
			},
			BlockRange: BlockRangeConf{
				From: 1,
				To:   -1,
			},
			Database: DatabaseConf{
				TypeName: "TIDB",
				User:     "root",
				Password: "123456",
				Endpoint: "127.0.0.1:4000",
			},
		}, jobs[0])
	})
}
