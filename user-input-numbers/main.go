/*
Create a program that prints to the terminal asking for a user to enter a small number and a larger number.
Print the remainder of the larger number divided by the smaller number.
*/

package main

import "fmt"

func main() {
	var smallNumber int
	var bigNumber int
	fmt.Print("Enter a small number: ")
	fmt.Scan(&smallNumber)
	fmt.Print("Enter a big number: ")
	fmt.Scan(&bigNumber)
	fmt.Println(smallNumber, bigNumber)
}
