package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

// It prints 4 cause by the time goroutine got the chance to run
// the value of i had already been incremented to value 4. To fix
// that, we need to pass the i as a parameter to the goroutine func
// so that goroutine operates on the input that has been passed to it
//
// Goroutines operate of the current value of the variable at the
// time of execution
//
// If we want the goroutines to operate on a specific value, then we
// need to pass that as an input to the goroutine.
func main() {
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
