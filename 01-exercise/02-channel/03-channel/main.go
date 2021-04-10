package main

import "fmt"

func main() {
	ch := make(chan int, 6)

	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			fmt.Printf("sending value %v\n", i)
			ch <- i
		}
	}()

	for v := range ch {
		fmt.Printf("received value %v\n", v)
	}
}
