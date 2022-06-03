package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

type Human struct {
	Age      uint
	Children uint
}

type Employee struct {
	// Similar to multi inheritance
	Person // Person embedded in Employee
	Human
	Salary float64
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

func (p Employee) Greet() {
	fmt.Println("Hello from Employee")
}

func NewEmployee(name string, age uint, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}
}

func NewHuman(age, children uint) Human {
	return Human(age, children)
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {
	e := NewEmployee("Alberto", 25, 2, 1000000)
	// fmt.Println(e.Name, e.Age) // conflict
	fmt.Println(e.Name, e.Human.Age)
	e.Greet()        // access the Greet method of Employee
	e.Person.Greet() // access the Greet method of Person
	e.Payroll()

}
