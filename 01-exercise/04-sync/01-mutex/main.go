package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	balance int
	wg      sync.WaitGroup
	mu      sync.Mutex
)

func main() {
	runtime.GOMAXPROCS(4)

	deposit := func(amount int) {
		mu.Lock()
		balance += amount
		mu.Unlock()
	}

	withdraw := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdraw(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
