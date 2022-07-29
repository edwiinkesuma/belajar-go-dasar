package main

import "fmt"

func factorialLoop(value int)int  {
	res := 1
	for i := value; i > 0; i-- {
		res *= i
	}
	return res
}

func factorialRecursive(value int)int{
	if value == 1 {
		return 1
	} else{
		return value * factorialRecursive(value -1)
	}
}

func main()  {
	fmt.Println(factorialLoop(5))
	fmt.Println(5*4*3*2*1)
	fmt.Println(factorialRecursive(5))

}