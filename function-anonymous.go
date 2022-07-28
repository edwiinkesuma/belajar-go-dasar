package main

import "fmt"

type Filter func(string)bool
func registerUser(name string, filter Filter)  {
	if filter(name) {
		fmt.Println("Welcome", name)
	} else{
		fmt.Println("You are blocked")
	}
}

func main()  {
	filtered := func(name string)bool{
		return name == "Edwin"
	}
	registerUser("Edwin", filtered)

	registerUser("Kesuma", func(name string)bool{ return name=="Kesuma"})
}