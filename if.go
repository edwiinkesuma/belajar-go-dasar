package main

import "fmt"

func main()  {
	name := "Edwin"

	if name == "Edwin"{
		fmt.Println("Hi Edwin")
	} else if name == "Kesuma"{
		fmt.Println("Hi Kesuma")
	} else {
		fmt.Println("Who are you?")
	}

	//if dengan short statement
	if length := len(name); length <= 5  {
		fmt.Println("I like ur name")
	} else{
		fmt.Println("Your name is too long")
	}
}