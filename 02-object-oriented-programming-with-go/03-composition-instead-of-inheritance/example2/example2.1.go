package main

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// alias declaration
type myAlias = course

// type definition, based on composite type (structure)
type newCourse course

func main() {
	c := course{name: "Go"}
	c.Print()
	fmt.Printf("The type is: %T\n", c)

	// myAlias refers to the course type
	c1 := myAlias{name: "Go"}
	c1.Print()
	fmt.Printf("The type is: %T\n", c1) // prints "The type is: main.course"
	
	// newcourse is a new type different from type course, they have the same fields
	c2 := newCourse{name: "Go"}
	/*c2.Print()*/ // It causes an error, because it does not inherit the Print method
	fmt.Printf("The type is: %T\n", c2) // prints "Type is: newCourse"
}