package main

import (
	"example.com/struct/user"
	"fmt"
	"strconv"
)

func main() {
	firstsName := getUserData("Introduce your first name: ")
	lastName := getUserData("Introduce your last name: ")
	birthday := getUserData("Introduce your birthday: ")
	per, err := user.New(firstsName, lastName, birthday)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	per.PrintPerson()
	roomNumber, _ := strconv.Atoi(getUserData("Introduce your room number: "))
	patient, err := user.NewPatient(firstsName, lastName, birthday, roomNumber)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	patient.PrintPatient()

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
