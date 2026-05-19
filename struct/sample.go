package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Method
func (customer Customer) sayHello() {
	fmt.Println("Hello", customer.Name)
}

func main() {
	var customer Customer
	customer.Name = "Budi"
	customer.Address = "Jakarta"
	customer.Age = 25

	customer.sayHello()

	customer2 := Customer{
		Name:    "Joko",
		Address: "Bandung",
		Age:     30,
	}

	customer2.sayHello()
}
