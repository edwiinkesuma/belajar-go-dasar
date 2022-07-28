package main

import "fmt"

func getPerson()(name string, age int){
	name = "Edwin"
	age = 22
	return
}

func main()  {
	a,b := getPerson()
	fmt.Println(a)
	fmt.Println(b)
}