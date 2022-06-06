package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

type GreeterByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

type Text string

func (p Person) Greet() {
	fmt.Printf("(person) Hi im %s\n", p.Name)
}

func (t Text) Greet() {
	fmt.Printf("(text) Hi im %s\n", t)
}

func (p Person) Bye() {
	fmt.Printf("Goodbye I'm %s\n", p.Name)
}

func (t Text) Bye() {
	fmt.Printf("Goodbye I'm %s\n", t)
}

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

func (p Person) String() string {
	return "Hi I'm a person and my name is: " + p.Name
}

func main() {
	p := Person{Name: "Albert"}
	fmt.Println(p)
}