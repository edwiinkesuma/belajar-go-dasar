package main

import "fmt"

func main()  {
	name := "Edwin"
	switch name {
	case "Edwin":
		fmt.Println("Hallo Edwin")
	case "Kesuma":
		fmt.Println("Hallo Edwin")
	default:
		fmt.Println("Who are you?")
	}

	//switch dengan short statement
	switch length := len(name); length < 5 {
	case true:
		fmt.Println("Your name is too long")
	default:
		fmt.Println("I like your name")
	}

	length2 := len(name)
	//Switch tanpa kondisi
	switch {
	case length2 > 5:
		fmt.Println("Your name is too long")
	default:
		fmt.Println("I like your name")
	}
}