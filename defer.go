package main

import "fmt"

func logging(){
	fmt.Println("Fungsi logging dipanggil")
}

func runApp(value int){
	defer logging() // defer digunakan untuk memanggil function diakhir function yang memanggil walaupun di function pemanggil ada error
	fmt.Println("Run app")
	res := 10 / value
	fmt.Println(res)
}

func main()  {
	runApp(0)
}