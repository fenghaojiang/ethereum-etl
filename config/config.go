package config

import (
	"encoding/json"
	"os"

	"github.com/fenghaojiang/ethereum-etl/common/log"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

type JobConf struct {
	JobID        uint           `json:"jobID"`
	Network      string         `json:"network"`
	ETLJobType   string         `json:"ETLJobType"`
	RPCEndpoints []string       `json:"RPCEndpoints"`
	BlockRange   BlockRangeConf `json:"blockRange"`
	Database     DatabaseConf   `json:"database"`
}

type BlockRangeConf struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type DatabaseConf struct {
	TypeName string `json:"type"`
	User     string `json:"user"`
	Password string `json:"password"`
	Endpoint string `json:"endpoint"`
}

func LoadConfig(path string) ([]JobConf, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Error("failed to read file", zap.String("path", path), zap.Error(err))
		return nil, err
	}
	jobContent := gjson.GetBytes(content, "jobs").String()
	var jobs []JobConf
	err = json.Unmarshal([]byte(jobContent), &jobs)
	if err != nil {
		log.Error("failed to unmarshal json to job conf",
			zap.String("json content", string(content)),
			zap.Error(err))
		return nil, err
	}
	return jobs, nil
}
