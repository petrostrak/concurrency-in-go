package main

import (
	"fmt"
	"sync"
)

var (
	sharedRsc = make(map[string]string)
	wg        sync.WaitGroup
)

func main() {
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
			// time.Sleep(1 * time.Millisecond)
		}

		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	c.L.Unlock()

	wg.Wait()
}
