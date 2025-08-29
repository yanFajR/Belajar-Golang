package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Jakarta"
	eko.Age = 30

	fmt.Println(eko)
	fmt.Println(eko.Name)
	eko.sayHello("Budi")

	joko := Customer {
		"Joko", "Bekasi", 30,
	}

	fmt.Println(joko)
	fmt.Println(joko.Name)
	joko.sayHello("Budi")
}
