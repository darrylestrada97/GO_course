package main

import (
	"fmt"
	"math"
)

// This application will ask the user for an amount of money, and will give the revenue of the money based on the
// return rate and the years of the investment.
func main() {
	const inflationRate = 2.5
	investmentAmount := float64(1000)
	expectedReturnRate := 5.5
	years := float64(10)

	fmt.Print("Introduce the invested value: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Introduce the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Introduce the number of years: ")
	fmt.Scan(&years)
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("Future value:", futureValue)
	fmt.Println("Future real value:", futureRealValue)
}
