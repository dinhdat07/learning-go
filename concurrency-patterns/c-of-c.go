package main

import "fmt"

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func client(f func([]int) int, clientRequests chan *Request) {
	request := &Request{[]int{3, 4, 5}, f, make(chan int)}
	clientRequests <- request
	fmt.Println("answer: %d\n", <-request.resultChan)
}

func handle(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}
