package common

type IRunnable interface {
	Start() error
	Stop() error
	Restart() error
}
