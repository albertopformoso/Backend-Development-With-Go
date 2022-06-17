package main

import "log"

func main() {
	out := Add(2, 3)
	log.Printf("Output: %d", out)

	out2 := AddAccumulative(1, 2, 3, 4, 5)
	log.Printf("Output 2: %d", out2)
}
