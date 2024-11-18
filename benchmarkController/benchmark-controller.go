package benchmarkController

import (
	"log/slog"
	"sync"
	"time"

	"github.com/Mtze/CI-Benchmarker/executor"
	"github.com/Mtze/CI-Benchmarker/persister"
)

// Benchmark represents a benchmarking process that includes an executor to run the benchmarks,
// a persister to save the results, and a job counter to keep track of the number of jobs executed.
type Benchmark struct {
	Executor   executor.Executor
	Persister  persister.Persister
	JobCounter int
}

// run executes the benchmark jobs concurrently. It logs the start of job execution,
// schedules each job, executes it using the provided executor, and stores the job
// result using the persister. It waits for all jobs to complete before returning.
//
// The function logs various stages of job execution, including the start of job
// scheduling, any errors encountered during execution, and the successful storage
// of job results.
func (b Benchmark) run() {
	slog.Debug("Running jobs", slog.Any("number", b.JobCounter), slog.Any("executor", b.Executor))
	var wg sync.WaitGroup

	for i := 0; i < b.JobCounter; i++ {
		wg.Add(1)

		go func(p persister.Persister) {
			defer wg.Done()
			slog.Debug("Scheduling job %d", slog.Any("i", i))
			// Execute the job
			uuid, err := b.Executor.Execute()
			if err != nil {
				slog.Error("Error while scheduling", slog.Any("error", err))
			}

			// Store the job
			slog.Debug("Storing job", slog.Any("uuid", uuid))
			p.StoreJob(uuid, time.Now())

			slog.Debug("Job send successfully", slog.Any("uuid", uuid))
		}(b.Persister)
	}

	wg.Wait()
}