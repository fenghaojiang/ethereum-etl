package common

import (
	"sync"

	"github.com/fenghaojiang/ethereum-etl/common/log"
	"go.uber.org/zap"
)

type Job struct {
	IRunnable
	JobUniqueKey string
}

type JobsManager struct {
	Jobs map[string]Job
}

func (jm *JobsManager) StartAllJobs() {
	var wg sync.WaitGroup
	wg.Add(len(jm.Jobs))
	for uniqueKey := range jm.Jobs {
		go func(key string) {
			defer wg.Done()
			jm.Jobs[key].Start()
			log.Info("job started", zap.String("jobID", key))
		}(uniqueKey)
	}
	wg.Wait()
	log.Info("all jobs started", zap.Int("length of jobs", len(jm.Jobs)))
}

func (jm *JobsManager) StopAllJobs() {
	for uniqueKey := range jm.Jobs {
		jm.Jobs[uniqueKey].Stop()
		log.Info("job stopped", zap.String("jobID", uniqueKey))
	}
	log.Info("All jobs stopped", zap.Int("length of jobs", len(jm.Jobs)))
}
