package main

import "fmt"

func main() {
	number := make(chan int, 2)
	signal := make(chan struct{})

	go receive(signal, number)
	send(number)

	<-signal
}

func send(number chan<- int) {
	number <- 10
	number <- 8
	number <- 6

	close(number)
}

func receive(signal chan<- struct{}, number <-chan int) {
	for v := range number {
			fmt.Println(v)
	}

	signal <- struct{}{}
}