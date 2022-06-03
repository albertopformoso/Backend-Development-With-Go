// Setters and Getters
package main

import (
	"encapsulation/course"
	"fmt"
)

func main() {
	Go := course.New("Go from scratch", 0, false)
	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Intro",
		2: "Structs",
		3: "Maps",
	})
	Go.SetName("OOP with Go")

	fmt.Printf("%+v\n", Go)
	fmt.Println(Go.Name())
	Go.PrintClasses()
}
