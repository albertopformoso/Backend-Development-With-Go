// Abstraction through structures
// // Methods With Pointer receivers
package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}

func (c Course) PrintClasses() {
	text := "The subjects are: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func main() {
	Go := Course{
		Name:    "Go from scratch",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Intro",
			2: "Structs",
			3: "Maps",
		},
	}

	
	(&Go).ChangePrice(67.12) // Also valid Go.ChangePrice(67.12)
	fmt.Printf("Course price: %v\n", Go.Price)
}
