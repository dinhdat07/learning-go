package main

import (
	"fmt"
	"sync"
)

// WorkerPool with defined type struct and implement with WaitGroup

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Square int
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		results <- Result{JobID: job.ID, Square: job.Value * job.Value}
	}
}

// receiver from results channel
func collectResults(results <-chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for result := range results {
		fmt.Printf("Job ID: %d, Input: %d, Squared Value: %d\n", result.JobID, result.JobID, result.Square)
	}
}

// orchestrates the entire operation
func dispatcher(jobCount, workerCount int) {
	jobs := make(chan Job, jobCount)
	results := make(chan Result, jobCount)

	// use wg to ensure all goroutine finish before this main goroutine
	var wg sync.WaitGroup

	wg.Add(workerCount)
	for w := 1; w <= workerCount; w++ {
		go worker(w, jobs, results, &wg) //pointer ref
	}

	var resultsWg sync.WaitGroup
	resultsWg.Add(1)
	go collectResults(results, &resultsWg)

	for j := 1; j <= jobCount; j++ {
		jobs <- Job{ID: j, Value: j}
	}
	close(jobs)
	wg.Wait()

	close(results)
	resultsWg.Wait()
}

// func main() {
// 	const jobCount = 100    // total number of jobs to process
// 	const workerCount = 3  // number of workers to process the jobs

// 	fmt.Println("Starting batch processing with synchronized result collection...")
// 	dispatcher(jobCount, workerCount)
// }
