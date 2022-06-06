package main

import "fmt"

type exampler interface {
	x()
}

func wrapper(i interface{}) {
	fmt.Printf("Value: %v, Type %T\n", i, i)
}

func main() {
	// var e exampler
	// fmt.Printf("Value: %v, Type %T", e, e)
	// e.x()

	wrapper(12)
	wrapper(14.32)
	wrapper(false)
	wrapper("Albert")
}