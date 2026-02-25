package main

import (
	"fmt"
)

// generate a source channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// func to run as multiple workers
func fanoutWorker(id int, in <-chan int) {
	for n := range in {
		fmt.Println("Worker", id, "got", n)
	}
}

// func main() {
// 	c := generator(1, 2, 3, 4, 5, 6)

// 	go worker(1, c)
// 	go worker(2, c)
// 	go worker(3, c)

// 	time.Sleep(time.Second)
// }
