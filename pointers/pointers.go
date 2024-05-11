package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("age before:", age)
	editAdultYears(agePointer)
	fmt.Println("adult years:", age)
}

func editAdultYears(age *int) {
	*age = *age - 18
}
