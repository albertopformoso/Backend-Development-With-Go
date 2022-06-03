package courses

import "fmt"

type Course struct { // <- Export Identifier is when you have the first letter uppercased
	Name    string // <- Export Identifier
	Price   float64 // <- Export Identifier
	IsFree  bool // <- Export Identifier
	UserIDs []uint // <- Export Identifier
	Classes map[uint]string // <- Export Identifier
}

func (c *Course) changePrice(price float64) { 
	c.Price = price
}

func (c *Course) PrintClasses() { // <- Export Identifier is when you have the first letter uppercased
	text := "The subjects are: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
