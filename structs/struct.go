package main

import (
	"fmt"
	"time"
)

type person struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func main() {
	firstsName := getUserData("Introduce your first name: ")
	lastName := getUserData("Introduce your last name: ")
	birthday := getUserData("Introduce your birthday: ")
	per := person{
		firstName: firstsName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}
	per.printPerson()
}
func (p person) getFullName() string {
	return p.firstName + " " + p.lastName
}

func (p person) printPerson() {
	fmt.Println("First name:", p.firstName)
	fmt.Println("Last name:", p.lastName)
	fmt.Println("Birthday:", p.birthday)
	fmt.Println("Created at:", p.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		return ""
	}
	return value
}
