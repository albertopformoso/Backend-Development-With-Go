// Abstraction through structures
package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
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

	css := Course{
		Name:   "CSS from scratch",
		IsFree: true,
	}

	js := Course{}
	js.Name = "JS from scratch"
	js.UserIDs = []uint{12, 67}

	fmt.Printf("Course Name: %v\n", Go.Name)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v\n", js)
}
