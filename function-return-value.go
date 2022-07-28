package main

import "fmt"


func main()  {
	fmt.Println(sayHello(""))
}

func sayHello(name string) string{
	if name == "" {
		return "Hi you, what is your name?"
	} else{
		return "Hello " + name
	}
}