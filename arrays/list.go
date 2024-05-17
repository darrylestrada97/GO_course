package main

import "fmt"

func main() {
	// Replace

	var productNames []string
	productNames = append(productNames, "iPhone")
	fmt.Println(productNames[0])
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices[3])
}
