package main

import "fmt"

func main()  {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"Jully",
		"Augustus",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[4:7]

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//Merubah isi array akan otomatis merubah isi slice
	// months[5] = "MayToo"
	// fmt.Println(slice1)

	//Merubah isi slice akan merubah isi array juga
	// slice1[0] = "AprilThen"
	// fmt.Println(months)

	slice2 := months[11:]
	fmt.Println(slice2)
	fmt.Println(cap(slice2))


	//jika append ke slice yang sudah punya kapasitas maksimum maka slice nya akan membuat array baru dan tidak mempengaruhi array lama
	slice3 := append(slice2, "newDecember")
	fmt.Println(slice3)
	slice3[0] = "DecemberToo"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	slice4 := months[3:4]
	fmt.Println("slice4")
	fmt.Println(slice4)
	fmt.Println(cap(slice4))

	slice5 := append(slice4, "newMonth")
	fmt.Println(slice5)
	fmt.Println(months)
	slice5[0] = "Hey"
	fmt.Println(slice5)
	fmt.Println(months)
	fmt.Println(slice4)


	fmt.Println("====================")
	newSlice := make([]string, 2, 5)
	fmt.Println(newSlice)
	newSlice[0] = "Edwin"
	newSlice[1] = "Kesuma"
	fmt.Println(newSlice)
	newSlice = append(newSlice, "Raja")
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("=====================")
	copySLice := make([]string, 3, 5)
	fmt.Println(copySLice)
	copy(copySLice, newSlice)
	fmt.Println(copySLice)

	fmt.Println("=====================")
	isArray := [...]int{1, 2, 3, 4, 5} //Array
	isSlice := []int{1, 2, 3, 4, 5 } //Slice

	fmt.Println(isArray)
	fmt.Println(isSlice)


}