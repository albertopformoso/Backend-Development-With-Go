package main

import "fmt"

// type definition, based on pre-declared type (bool)
type newBool bool

func (b newBool) String() string {
	if b {
		return "TRUE"	
	}
	return "FALSE"
}

func main() {
	var b newBool = true
	fmt.Println(b.String()) // prints "TRUE"
}