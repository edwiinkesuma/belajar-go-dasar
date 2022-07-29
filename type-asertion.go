package main

import "fmt"

func random()interface{}  {
	return "Hallo"
}

func main()  {
	res := random()

	switch value := res.(type) {
	case string:
		fmt.Println(value, "adalah string")
	case int:
		fmt.Println(value, "adalah int")
	default:
		fmt.Println("Unkown")
	}
}