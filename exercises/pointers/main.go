package main

import "fmt"

func main() {
	age := 32
	agePointer := &age

	fmt.Println("Age:", *agePointer)

	editToAdultYears(agePointer)

	fmt.Println(age)
}

func editToAdultYears(age *int) {
	*age = *age - 18
}
