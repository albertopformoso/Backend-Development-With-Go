// Encapsulation
package main

import (
	"encapsulation/courses"
)

func main() {
	Go := &courses.Course{
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

	Go.PrintClasses()
	Go.changePrice(24)
}