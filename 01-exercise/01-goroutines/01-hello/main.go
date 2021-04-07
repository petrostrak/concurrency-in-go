package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("This is a direct call")

	// goroutine function call
	go fun("Goroutine-1")

	// goroutine with anonymous func
	go func() {
		fun("Goroutine-2")
	}()

	// goroutine with function value call
	fv := fun
	go fv("Goroutine-3")

	// wait for goroutines to return
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done..")
}
