package service

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/fenghaojiang/ethereum-etl/model"
	"github.com/fenghaojiang/ethereum-etl/service"
)

var _ service.IJob = (*TransactionJob)(nil)

type TransactionJob struct {
	jobRunning int32
	batchSize  int
	ctx        context.Context
}

func NewTransactionJob(ctx context.Context) service.IJob {
	return &TransactionJob{}
}

func (j *TransactionJob) Start() error {
	if atomic.LoadInt32(&j.jobRunning) == service.JobStop {
		go j.Run()
		atomic.StoreInt32(&j.jobRunning, service.JobRunning)
		return nil
	}
	return fmt.Errorf("job started once")
}

func (j *TransactionJob) Stop() error {
	if atomic.LoadInt32(&j.jobRunning) == service.JobStop {
		return fmt.Errorf("job already stopped once")
	}

	// TODO Job Stop Job
	atomic.StoreInt32(&j.jobRunning, service.JobStop)
	return nil
}

func (j *TransactionJob) Run() {

	// fetch rpc
	resultChan := make(chan []model.Transaction, j.batchSize)

	go j.rpcRequest(resultChan)

	for {
		select {
		case _, ok := <-resultChan:
			if !ok {
				return
			}

			//TODO batch write txns
		case <-j.ctx.Done():

			return
		}
	}
}

func (j *TransactionJob) rpcRequest(outChan chan<- []model.Transaction) {

}
