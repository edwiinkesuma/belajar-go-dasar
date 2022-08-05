package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	address1 := Address {"Bandung", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1
	address4 := &Address{"Surabaya", "Jatim", "Indonesia"}
	address5 := new(Address)

	//tanda * digunakan untuk mengubah semua value di pointer yang sama
	*address2 = Address{"Malang", "Jawa TImur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
	fmt.Println(address5)

}