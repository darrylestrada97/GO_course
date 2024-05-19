package main

import "fmt"

type strString map[string]string

func (f strString) print() {
	for key, value := range f {
		fmt.Printf("%s: %s\n", key, value)
	}
}

func main() {
	// Replace

	userNames := make(strString, 10)
	fmt.Printf("Address of userNames map: %p\n", &userNames)
	fmt.Printf("Type of userNames map: %T\n", userNames)
	fmt.Println(userNames, len(userNames))
	userNames["Roy"] = "Raj"
	userNames["Ford"] = "John"
	userNames["Amy"] = "Winehouse"

	fmt.Println(userNames, len(userNames))
	delete(userNames, "Ford")
	fmt.Println("After deleting Ford")
	fmt.Println(userNames, len(userNames))

	userNames.print()

}
