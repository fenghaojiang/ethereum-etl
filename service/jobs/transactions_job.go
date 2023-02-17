package service

import (
	"fmt"
	"sync/atomic"

	"github.com/fenghaojiang/ethereum-etl/service"
)

type TransactionJob struct {
	jobRunning int32
}

func (j *TransactionJob) Start() error {
	if atomic.LoadInt32(&j.jobRunning) == service.JobStop {

		//TODO Start Job
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
