package main

import "fmt"

func sumAll(numbers ...int) int{
	result := 0
	for _, number := range numbers{
		result += number
	}
	return result
}

func main(){
	total := sumAll(10,10,10)
	fmt.Println(total)

	num := []int{20,20,20}
	total2 := sumAll(num...)
	fmt.Println(total2)
}