package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

// The function returns but goroutine still has the access to the
// local variable of the function.
//
// Usually when the function returns, the local variables go out
// out scope. But here, the runtime is clever enough to see that
// the reference to a local variable i is still being held by the
// goroutine, so it pins it. It moves it from the stack to heap,
// so that goroutine still has the access to the variable even
// after the enclosing function returns.
func main() {
	incr := func(wg *sync.WaitGroup) {
		var i int
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value of i: %v\n", i)
		}()
		fmt.Println("return from function")
		return
	}
	incr(&wg)
	wg.Wait()
	fmt.Println("Done..")
}
