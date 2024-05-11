package main

import "fmt"

func main() {
	firstsName := getUserData("Introduce your first name: ")
	lastName := getUserData("Introduce your last name: ")
	birthday := getUserData("Introduce your birthday: ")

	fmt.Println(firstsName, lastName, birthday)
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
