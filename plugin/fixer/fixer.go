package fixer

var _ IFixer = (*Fixer)(nil)

type IFixer interface {
	CheckReorg(blockNumber uint64, currentBlockHash string) (previousBlockHash string, isReorg bool)
	fixReorg(ch chan<- struct{}) error
}

type Fixer struct {
}

func (f *Fixer) CheckReorg(blockNumber uint64, currentBlockHash string) (previousBlockHash string, isReorg bool) {
	return "", false
}

func (f *Fixer) fixReorg(ch chan<- struct{}) error {
	return nil
}
