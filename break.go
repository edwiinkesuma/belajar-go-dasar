package main

import "fmt"

func main()  {
	for i:=0; i <=10; i++{
		fmt.Println("Perulangan ke ", i)

		if i == 6 {
			break
		}
	}
}