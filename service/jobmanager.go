package service

import (
	"fmt"
	"sync"

	"github.com/fenghaojiang/ethereum-etl/common/log"
	"go.uber.org/zap"
)

type JobsManager struct {
	Jobs sync.Map
}

func (jm *JobsManager) StartAllJobs() {
	var wg sync.WaitGroup
	jm.Jobs.Range(func(jobID, jobs any) bool {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()

			jobList := jobs.([]IJob)

			for i := range jobList {
				err := jobList[i].Start()
				if err != nil {
					log.Error("failed to start job", zap.String("jobID", key), zap.Error(err))
					continue
				}
			}
			log.Info("job started", zap.String("jobID", key))
		}(jobID.(string))
		return true
	})
	wg.Wait()
	log.Info("all jobs started")
}

func (jm *JobsManager) StopAllJobs() {
	jm.Jobs.Range(func(jobID, job any) bool {
		err := job.(IJob).Stop()
		if err != nil {
			log.Error("failed to stop job", zap.Any("jobID", jobID), zap.Error(err))
		}
		return true
	})
	log.Info("All jobs stopped")
}

func (jm *JobsManager) StopJob(jobID string) error {
	job, ok := jm.Jobs.Load(jobID)
	if !ok {
		return fmt.Errorf("job not exist")
	}
	return job.(IJob).Stop()
}

func (jm *JobsManager) StartJob(jobID string) error {
	job, ok := jm.Jobs.Load(jobID)
	if !ok {
		return fmt.Errorf("job not exist")
	}
	return job.(IJob).Start()
}

func (jm *JobsManager) AddJob(jobID string, job IJob) {
	jm.Jobs.Store(jobID, job)
}
