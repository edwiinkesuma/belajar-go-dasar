package main

import "fmt"

type Customer struct{
	Name, Address string
	Age int
}

func (customer Customer) sayHi()  {
	fmt.Println("Hi", customer.Name)
}

func main()  {
	edwin := Customer{
		Name: "edwin",
		Address: "Madiun",
		Age: 22,
	}
	edwin.sayHi()
}