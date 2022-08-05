package main

import "fmt"

type Man struct{
	Name string
}

func (man *Man) getMaried()  {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{"Edwin"}
	fmt.Println(man.Name)
	man.getMaried()
	fmt.Println(man.Name)
}