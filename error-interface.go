package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int)(int, error)  {
	if pembagi == 0 {
		return 0, errors.New("NIlai pembagi tidak boleh Nol") 
	} else {
		return nilai/pembagi, nil
	}
}

func main()  {
	nilai, eror := pembagi(10, 0)

	if eror == nil {
		fmt.Println(nilai)	
	}else{
		fmt.Println(eror)
	}
}