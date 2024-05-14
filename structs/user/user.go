package user

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func (p Person) getFullName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) PrintPerson() {
	fmt.Println("First name:", p.firstName)
	fmt.Println("Last name:", p.lastName)
	fmt.Println("Birthday:", p.birthday)
	// format createdAt time to human-readable format
	fmt.Println("created at:", p.createdAt.Format("2006-01-02 15:04:05"))

}

func NewPerson(firstName, lastName, birthday string) (*Person, error) {

	if firstName == "" || lastName == "" || birthday == "" {
		return nil, errors.New("all fields are required")
	}
	return &Person{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}, nil
}
