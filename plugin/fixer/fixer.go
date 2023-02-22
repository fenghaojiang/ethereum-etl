package fixer

type IFixer interface {
	CheckReorg(blockNumber uint64, currentBlockHash string) (previousBlockHash string, isReorg bool)
	fixReorg(ch chan<- struct{}) error
}
