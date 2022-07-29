package main

import "fmt"

func endApp()  {
	message := recover()

	if message != nil {
		fmt.Println("Error: ", message)
	}

	fmt.Println("Aplikasi berhenti")
}

func startApp(error bool)  {
	defer endApp()
	fmt.Println("Mulai aplikasi")
	if error {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main()  {
	startApp(true)
}