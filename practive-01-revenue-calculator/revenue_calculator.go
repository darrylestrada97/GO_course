package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue= scanValue("Introduce the revenue:")
	expenses= scanValue("Introduce the expenses:")
	taxRate= scanValue("Introduce the tax rate:")
	// fmt.Print("Introduce the revenue: ")
	// fmt.Scan(&revenue)
	// fmt.Print("Introduce the expenses: ")
	// fmt.Scan(&expenses)
	// fmt.Print("The tax rate of the revenue calculator: ")
	// fmt.Scan(&taxRate)

	earningsBT, earningsAT, ratio := calculateFutureValues(revenue, expenses, taxRate)

	fmt.Printf("Earning before taxes: %v\n", earningsBT)
	fmt.Printf("Earning after taxes: %v\n", earningsAT)
	fmt.Printf("ratio: %.2f", ratio)

}
func calculateFutureValues(revenue, expenses, taxRate float64) (earningsBT float64, earningsAT float64, ratio float64) {
	earningsBT = revenue - expenses
	earningsAT = earningsBT * (1 - taxRate/100)
	ratio = earningsBT / earningsAT
	return
}

//function that get a text and prints the text and scans the value from the user, return the value the user inputs.
func scanValue(text string) (value float64) {

	fmt.Print(text)
	fmt.Scan(&value)
	return value
}
