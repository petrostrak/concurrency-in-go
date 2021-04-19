package main

import "fmt"

// TODO: Build a pipeline
// generator() -> square() -> print

// generator - convertes a list of ints to a channel
func generator(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()

	return out
}

// square - recv on inbound channel
// squares the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	// set up pipeline
	// ch := generator(2, 3, 4)
	// out := square(ch)

	for n := range square(generator(2, 3, 4)) {
		fmt.Println(n)
	}
}
