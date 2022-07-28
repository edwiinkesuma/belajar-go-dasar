package main

import "fmt"

func main()  {
	angka := 5
	for angka <=10 {
		fmt.Println("Hallo")
		angka++
	}

	for number := 0; number <=10; number++{
		fmt.Println("Hi")
	}

	newSlices := []string{"Edwin", "Kesuma"}

	for _, slice := range newSlices {
		fmt.Println(slice)
	}

	person := make(map[string]string)
	person["name"] = "Edwin"
	person["job"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}