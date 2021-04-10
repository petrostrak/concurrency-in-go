package main

import "fmt"

func genMsg(ch1 chan<- string) {
	// send msg on ch1
	ch1 <- "Sending msg"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	// receive msg on ch1
	// send in on ch2
	in := <-ch1
	ch2 <- in
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	fmt.Println(<-ch2)
}
