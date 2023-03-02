package service

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fenghaojiang/ethereum-etl/common/log"
)

const (
	TransactionsJobType JobType = "Transactions"
	LogsJobType         JobType = "Logs"
	ERC20BalanceType    JobType = "ERC20Balance"
)

type JobType string

type IJob interface {
	Start() error
	Stop() error
	Done() <-chan struct{}
}

const (
	JobStop    int32 = 0
	JobRunning int32 = 1
)

func Run(j IJob) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	for {
		go j.Start()
		var s os.Signal
		var done bool
		select {
		case s = <-signalChan:

		case <-j.Done():
			s = syscall.SIGQUIT
			done = true
		}
		switch s {
		case syscall.SIGINT:
			log.Warn("event:service msg:Process exit with SIGINT!")
		case syscall.SIGQUIT:
			if done {
				log.Warn("event:service Process exit with Runable Done")
			} else {
				log.Warn("event:service process exit with SIGQUIT!")
			}
		case syscall.SIGHUP:
			time.Sleep(time.Second)
			j.Stop()
			log.Warn("event:service process restart with SIGHUP...")
			continue // continue可以重启服务
		case syscall.SIGKILL:
			log.Warn("event:service process killed with SIGKILL!")
		case syscall.SIGTERM:
			log.Warn("event:service process exit with SIGTERM!")
		default:
			log.Warn("event:service process unknown exit")
		}

		// receive the stop signal
		j.Stop()

		break
	}
}
