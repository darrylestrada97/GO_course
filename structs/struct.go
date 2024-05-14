package main

import (
	"example.com/struct/user"
	"fmt"
)

func main() {
	firstsName := getUserData("Introduce your first name: ")
	lastName := getUserData("Introduce your last name: ")
	birthday := getUserData("Introduce your birthday: ")
	per, err := user.NewPerson(firstsName, lastName, birthday)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	per.PrintPerson()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, err := fmt.Scanln(&value)
	if err != nil {
		return ""
	}
	return value
}
