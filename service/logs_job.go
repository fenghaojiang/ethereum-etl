package service

import (
	"fmt"
	"sync/atomic"
)

type LogsJob struct {
	jobRunning int32
}

func (l *LogsJob) Start() error {
	if atomic.LoadInt32(&l.jobRunning) == 0 {

		atomic.StoreInt32(&l.jobRunning, 1)
		return nil
	}
	return fmt.Errorf("job started once")
}
