package main

import "fmt"

type Nama interface{
	getName()string
}

func sayHello(nama Nama){
	fmt.Println("Hallo", nama.getName())
}

type Person struct{
	Name string
}

func (person Person) getName()string {
	return person.Name
}

func main()  {
	edwin := Person{
		Name: "Edwin",
	}	

	sayHello(edwin)
}