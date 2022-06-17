package main

import "log"

func main() {
	out := fibonacciFor(10)
	log.Printf("Output: %d", out)

	out2 := fibonacciRec(10)
	log.Printf("Output 2: %d", out2)
	
	out3 := fibonacciRecMem(10)
	log.Printf("Output 3: %d", out3)
}
