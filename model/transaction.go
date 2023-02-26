package model

type Transaction struct {
	UniqueKey       string `gorm:"column:uniqueKey"`
	BlockNumber     uint64 `gorm:"column:blockNumber"`
	Timestamp       int64  `gorm:"column:timestamp"`
	BlockHash       string `gorm:"column:blockHash"`
	TxHash          string `gorm:"column:txHash"`
	ContractAddress string `gorm:"column:contractAddress"`
	From            string `gorm:"column:from"`
	To              string `gorm:"column:to"`
	Value           string `gorm:"column:value"`
}
