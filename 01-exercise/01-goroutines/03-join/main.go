package main

import (
	"fmt"
	"sync"
)

var (
	data int
	wg   *sync.WaitGroup
)

func main() {
	wg.Add(1)
	go func() {
		data++
		wg.Done()
	}()
	wg.Wait()

	fmt.Printf("the value of the data is %v\n", data)
	fmt.Println("Done..")
}
