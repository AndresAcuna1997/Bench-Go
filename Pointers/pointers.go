package main

import "fmt"

func main() {
	age := 27
	agePointer := &age

	fmt.Println(age)

	//fmt.Print(agePointer)
	getAdultYears(agePointer)

	fmt.Println(age)

}

func getAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}
