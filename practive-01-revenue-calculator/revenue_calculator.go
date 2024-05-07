package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Introduce the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Introduce the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("The tax rate of the revenue calculator: ")
	fmt.Scan(&taxRate)

	earningsBT := revenue - expenses
	earningsAT := earningsBT * (1-taxRate / 100)
	ratio := earningsBT / earningsAT
	fmt.Println("Earning before taxes: ",earningsBT)
	fmt.Println("Earning after taxes: ",earningsAT)
	fmt.Println("ratio: ",ratio)

}
