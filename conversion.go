package main

import "fmt"
func main()  {
	var int32 int32 = 128
	var int64 int64 = int64(int32)
	var int8 int8 = int8(int64) //jika value variable lebih besar dari nilai maks yang bisa ditampung maka valuenya akan dimulain dari dinilai minimum

	fmt.Println(int32)
	fmt.Println(int64)
	fmt.Println(int8)

	name := "Edwin Kesuma"
	e := name[0]
	eString := string(e)

	fmt.Println(name)
	fmt.Println(eString)
}