package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(val int) int {
		return val * 2
	})
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	var dNumbers []int

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
