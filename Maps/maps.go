package main

import "fmt"

func main() {
	// Replace

	userNames := make(map[string]string, 10)
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

}
