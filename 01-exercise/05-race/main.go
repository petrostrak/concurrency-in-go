package main

import (
	"fmt"
	"math/rand"
	"time"
)

// go build -race main.go
func main() {
	start := time.Now()
	var t *time.Timer
	ch := make(chan bool)

	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		// t.Reset(randomDuration())
		ch <- true
	})

	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}

	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
