package main

import "fmt"

func main() {
	// Replace

	var productNames []string
	productNames = append(productNames, "iPhone")
	fmt.Println(productNames[0])
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	featuredPrices := prices[1:3]
	fmt.Println(prices[3])
	bestPrices := featuredPrices[:1]

	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(bestPrices)
	fmt.Println(len(bestPrices), cap(bestPrices))
	bestPrices = bestPrices[:3]
	fmt.Println(bestPrices)
	fmt.Println(len(bestPrices), cap(bestPrices))

	prices1 := []float64{10.99, 9.99, 45.99, 20.0}
	fmt.Printf("Address of prices1 slice: %p\n", &prices1)
	prices1 = append(prices1, 1.99)
	fmt.Printf("Address of prices1 slice: %p\n", &prices1)
}
