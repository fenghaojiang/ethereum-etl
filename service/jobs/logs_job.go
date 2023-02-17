package service

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/fenghaojiang/ethereum-etl/service"
)

var _ service.IJob = (*LogsJob)(nil)

type LogsJob struct {
	jobRunning int32
	ctx        context.Context
	batchSize  int32
	cursor     uint64
	limitChan  chan struct{}
}

func NewLogsJob(ctx context.Context) service.IJob {
	return &LogsJob{}
}

func (j *LogsJob) Start() error {
	if atomic.LoadInt32(&j.jobRunning) == service.JobStop {

		atomic.StoreInt32(&j.jobRunning, service.JobRunning)
		return nil
	}
	return fmt.Errorf("job started once")
}

func (j *LogsJob) Stop() error {
	if atomic.LoadInt32(&j.jobRunning) == service.JobStop {
		return fmt.Errorf("job already stopped once")
	}

	// TODO Job Stop Job
	atomic.StoreInt32(&j.jobRunning, service.JobStop)
	return nil
}
