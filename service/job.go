package service

const (
	TransactionsJobType JobType = "Transactions"
	LogsJobType         JobType = "Logs"
	ERC20BalanceType    JobType = "ERC20Balance"
)

type JobType string

type IJob interface {
	Start() error
	Stop() error
}

const (
	JobStop    int32 = 0
	JobRunning int32 = 1
)
