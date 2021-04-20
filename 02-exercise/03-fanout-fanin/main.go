package main

import (
	"fmt"
	"sync"
)

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

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := generator(2, 3, 4)

	ch1 := square(in)
	ch2 := square(in)
	ch3 := square(in)

	for n := range merge(ch1, ch2, ch3) {
		fmt.Println(n)
	}
}
