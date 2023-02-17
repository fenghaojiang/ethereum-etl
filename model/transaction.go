package model

type Transaction struct {
	BlockNumber     uint64
	Timestamp       int64
	BlockHash       string
	TxHash          string
	ContractAddress string
	From            string
	To              string
	Value           string
}
