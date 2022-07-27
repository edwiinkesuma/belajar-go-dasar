package main

import "fmt"
func main()  {
	//type digunakan untuk membuat alias atau mendeklarasikan tipe data baru darri tipe data yang sudah ada
	type newString string

	var name newString = "Edwin Kesuma"
	fmt.Println(name)
}