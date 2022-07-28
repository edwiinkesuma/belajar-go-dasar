package main

import "fmt"

func sayHi(name string, filter func(string)string)  {
	filtered := filter(name)

	fmt.Println(filtered)
}

func filteredName(name string) string {
	if name == "Anjing" {
		return "...."
	} else{
		return name;
	}
}

func main()  {
	filter := filteredName

	sayHi("Edwin", filter)
}