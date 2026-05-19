package main

import "fmt"

// Interface (Interface is a collection of method signatures)
type Person interface {
	GetName() string
}

func sayHello(person Person) {
	fmt.Println("Hello", person.GetName())
}

type Employee struct {
	Name string
}

func (e Employee) GetName() string {
	return e.Name
}

func main() {
	// You'll need a struct that implements the Person interface to use here
	person := Employee{
		Name: "Budi",
	}

	sayHello(person)
}
