/*
Write a function with one variadic parameter that finds the greatest number in a list of numbers.
*/

package main

import "fmt"

func greatest(numbers ...int) int {
	var greatestNumber int
	for _, v := range numbers {
		if v > greatestNumber {
			greatestNumber = v
		}
	}
	return greatestNumber
}

func main() {
	fmt.Println(greatest(1, 2, 10, 3, 8, 6, 5))
}
