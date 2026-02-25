package main

import (
	"fmt"
	"time"
)

func basicWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker id:", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("Worker id:", id, "finished job", j)
		results <- (j * 2)
	}
}

// func main() {
// 	numJobs := 10
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)
// 	numWorkers := 3

// 	for i := range numWorkers {
// 		go basicWorker(i, jobs, results)
// 	}

// 	for j := range numJobs {
// 		jobs <- j
// 	}
// 	close(jobs) // close signal from sender, so worker know when to stop for-loop

// 	for range numJobs {
// 		<-results
// 	}

// }
