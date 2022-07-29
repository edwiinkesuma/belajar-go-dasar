package main

import "fmt"

func hello(a int)interface{}  {
	if a==1 {
		return 1
	} else{
		return "hi"
	}
}

func main()  {
	fmt.Println(hello(2))
}