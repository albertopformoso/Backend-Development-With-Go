package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func hello() {
	fmt.Println("Hello, Edteam Community")
	wg.Done()
}

func main() {
	wg.Add(2)
	go hello()
	go func(){
		fmt.Println("Hello, from anonymous function")
		wg.Done()
	}()

	fmt.Println("Hello, Gophers")
	wg.Wait()
}
