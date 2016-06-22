package main

import "fmt"

func main() {
	// normal for loop
	fmt.Println("Start for loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// nested for loop
	fmt.Println("\nStart nested loop")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}
