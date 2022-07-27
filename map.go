package main

import "fmt"

func main()  {
	person := map[string]string{
		"name" : "Edwin",
		"address" : "Madiun",
	}
	fmt.Println(person)

	person["title"] = "Programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Belajar GO"
	book["author"] = "Edwin"
	book["ups"] = "salah"
	fmt.Println(book)
	
	delete(book, "ups")
	fmt.Println(book)



}