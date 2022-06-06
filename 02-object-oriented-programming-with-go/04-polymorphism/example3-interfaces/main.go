package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
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

func GreetAll(gs ...Greeter) {
	for _, g := range gs {
		g.Greet()
		fmt.Printf("\t Value: %v, Type: %T\n", g, g)
	}
}

func ByeAll(bs ...Byer) {
	for _, b := range bs {
		b.Bye()
	}
}

func main() {
	p := Person{Name: "Albert"}
	var t Text = "Daisy"
	GreetAll(p, t)
	ByeAll(p, t)
}