package main

import "fmt"

func main()  {
	var names [2]string

	names[0] = "Edwin"
	names[1] = "Kesuma"

	fmt.Println(names)

	values := [3]int{
		90,
		85,
		80,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}