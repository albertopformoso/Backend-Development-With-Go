// Function New() instead of Constructor method
package main

import (
	"encapsulation/course"
	"fmt"
)

func main() {
	Go := course.New("Go from scratch", 0, false)
	Go.UserIDs = []uint{12, 56, 89}
	Go.Classes = map[uint]string{
			1: "Intro",
			2: "Structs",
			3: "Maps",
		}

	fmt.Printf("%+v\n", Go)
}