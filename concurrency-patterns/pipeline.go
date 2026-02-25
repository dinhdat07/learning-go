// package main

// import "fmt"

// func gen(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for _, n := range nums {
// 			out <- n
// 		}
// 	}()

// 	return out
// }

// func sq(in <-chan int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		defer close(out)
// 		for n := range in {
// 			out <- n * n
// 		}
// 	}()

// 	return out
// }

// func log(in <-chan int) {
// 	for n := range in {
// 		fmt.Println(n)
// 	}
// }

// // func main() {
// // 	// set up the pipeline with gen
// // 	c := gen(2, 5, 7, 8, 3)
// // 	out := sq(c)
// // 	log(out)
// // }
