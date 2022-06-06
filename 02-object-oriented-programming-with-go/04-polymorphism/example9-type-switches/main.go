package main

import (
	"fmt"
	"strings"
)

func wrapper(i interface{}) {
	// fmt.Printf("Value: %v, Type %T\n", i, i)

	switch v := i.(type) {
	case string:
		fmt.Printf("%v\n", strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Printf("Value: %v, Type %T\n", i, i)
	}
}

func main() {
	wrapper(12)
	wrapper(14.32)
	wrapper(false)
	wrapper("Albert")
}