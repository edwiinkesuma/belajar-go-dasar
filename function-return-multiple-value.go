package main

import "fmt"

func getFullName()(string, string)  {
	return "Edwin", "Kesuma"
}

func main()  {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	//gunakan tanda _ untuk ignore
	job,_ := getJob()
	fmt.Println(job)
}

func getJob()(string, string){
	return "Programmer", "Fulltime"
}