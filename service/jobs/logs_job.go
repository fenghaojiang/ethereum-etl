package service

import (
	"context"
	"fmt"
	"sync/atomic"

	"github.com/ethereum/go-ethereum"
	"github.com/fenghaojiang/ethereum-etl/service"
)

var _ service.IJob = (*LogsJob)(nil)

type LogsJob struct {
	jobRunning int32
	topics     ethereum.FilterQuery
}

func NewLogsJob(ctx context.Context, topics ethereum.FilterQuery) service.IJob {
	return &LogsJob{
		jobRunning: 0,
		topics:     topics,
	}
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
