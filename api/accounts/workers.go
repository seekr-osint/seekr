package accounts

import (
	"log"
	"sync"
)

func ToWorkerResult(scanResult *ScanResult, err error) *WorkerResult {
	return &WorkerResult{
		ScanResult: *scanResult,
		Error:      err,
	}
}

func Worker(id int, jobs <-chan Job, results chan<- WorkerResult) {
	log.Printf("Starting worker %d\n", id)
	for job := range jobs {
		res := *ToWorkerResult(job.RunScannerDefaultAccountResult())
		res.SetID(job.ID)
		results <- res
	}
}

func (j Jobs) StartWorkers(workers int) (*ScanResults, error) {
	jobs := make(chan Job, len(j))
	results := make(chan WorkerResult, len(j))
	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			Worker(workerID, jobs, results)
		}(i)
	}
	for _, job := range j {
		jobs <- job
	}
	close(jobs)
	wg.Wait()
	close(results)
	scanResults := ScanResults{}
	for result := range results {
		if result.Error != nil {
			log.Printf("error account scan:%s", result.Error)
			return nil, result.Error
		}
		scanResults = append(scanResults, result.ScanResult)
	}

	return &scanResults, nil
}
