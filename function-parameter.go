package main

import "fmt"

func sayHello(firstName string, lastName string){
	fmt.Println("Hallo", firstName, lastName)
}
func main()  {
	sayHello("Edwin", "Kesuma")
}