package course

import "fmt"

type course struct { 
	Name    string 
	Price   float64 
	IsFree  bool 
	UserIDs []uint 
	Classes map[uint]string 
}

func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name: name,
		Price: price,
		IsFree: isFree,
	}
}

func (c *course) changePrice(price float64) { 
	c.Price = price
}

func (c *course) PrintClasses() {
	text := "The subjects are: "
	for _, class := range c.Classes {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}
