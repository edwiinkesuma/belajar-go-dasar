package main

import "fmt"

func getName(name string)string{
	return "Hai " + name
}

func main(){
	stringGetName := getName
	fmt.Println(stringGetName("Edwin"))
}