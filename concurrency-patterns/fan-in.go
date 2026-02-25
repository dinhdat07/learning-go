package main

import "sync"

// variadic to receive from multiple channels
func merge(in ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(in))
	for _, c := range in {
		go output(c)
	}

	go func() {
		wg.Wait()
	}()
	// put Wait() in new goroutine to avoid block, this way will return the out channel immediately

	return out
}
