package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			// if out is block, and done was closed -> return, end this goroutine
			case <-done:
				return
			}
		}
	}()

	return out
}

func log(done chan struct{}, in <-chan int) {
	defer close(done)
	fmt.Println(<-in)
}

func main() {
	// set up the pipeline with gen
	c := gen(2, 5, 7, 8, 3)

	// create new channel done for cancellation
	done := make(chan struct{})
	out := sq(done, c)
	log(done, out)
}
