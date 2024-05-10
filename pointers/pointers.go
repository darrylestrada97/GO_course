package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("age before:", age)
	fmt.Println("adult years:", getAdultYears(agePointer))
}

func getAdultYears(age *int) int {
	return *age - 18
}
