package main

type Vector []float64

func (v Vector) Cal(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += v[i]
	}

	c <- 1
}

const numCPU = 4 // number of CPU cores

// parallel with fan-out pattern
func (v Vector) CalAll(u Vector) {
	c := make(chan int, numCPU) // Buffering optional but sensible.
	for i := 0; i < numCPU; i++ {
		go v.Cal(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
	}

	// Drain the channel.
	for i := 0; i < numCPU; i++ {
		<-c // wait for one task to complete
	}

}
