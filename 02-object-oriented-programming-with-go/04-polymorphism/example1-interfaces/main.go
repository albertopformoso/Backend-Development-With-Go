package main

import "fmt"

// Greeter is an interface
type Greeter interface {
	Greet()
}

// Person is a struct which implements the method Greet
type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("(person) Hi im %s\n", p.Name)
}

// Text is a defined type which implements the Greet method
type Text string

func (t Text) Greet() {
	fmt.Printf("(text) Hi im %s\n", t)
}

func main() {
	var p Greeter = Person{Name: "Albert"}
	p.Greet()

	var t Greeter = Text("Daisy") // casting since it is not a string
	t.Greet()
}