package service

import (
	"fmt"
	"sync/atomic"

	"github.com/fenghaojiang/ethereum-etl/service"
)

var _ service.IJob = (*LogsJob)(nil)

type LogsJob struct {
	jobRunning int32
}

func (l *LogsJob) Start() error {
	if atomic.LoadInt32(&l.jobRunning) == service.JobStop {

		//TODO Start Job
		atomic.StoreInt32(&l.jobRunning, service.JobRunning)
		return nil
	}
	return fmt.Errorf("job started once")
}

func (l *LogsJob) Stop() error {
	if atomic.LoadInt32(&l.jobRunning) == service.JobStop {
		return fmt.Errorf("job already stopped once")
	}

	// TODO Job Stop Job
	atomic.StoreInt32(&l.jobRunning, service.JobStop)
	return nil
}
