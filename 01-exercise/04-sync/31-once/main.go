package main

import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup
	once sync.Once
)

func main() {
	load := func() {
		fmt.Println("Run only once initialization function")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			once.Do(load)
		}()
	}
	wg.Wait()
}
