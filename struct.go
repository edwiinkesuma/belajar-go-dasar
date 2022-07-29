package main

import "fmt"

func main()  {
	type Customer struct{
		Name, Address string
		Age int
	}

	var edwin Customer
	edwin.Name = "Edwin"
	edwin.Address = "Madiun"
	edwin.Age = 22
	fmt.Println(edwin)
	fmt.Println(edwin.Name)
	fmt.Println(edwin.Age)
	fmt.Println(edwin.Address)

	kesuma := Customer{
		Name: "Kesuma",
		Address: "Madiun",
		Age: 22,
	}
	fmt.Println(kesuma)
	fmt.Println(kesuma.Name)
	fmt.Println(kesuma.Age)
	fmt.Println(kesuma.Address)

}