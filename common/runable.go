package common

type Runnable interface {
	Start() error
	Stop() error
	Restart() error
}
