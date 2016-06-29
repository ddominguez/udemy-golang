/*
Write a function which takes an integer. The function will have two returns.
The first return should be the argument divided by 2.
The second return should be a bool that let’s us know whether or not the argument was even.
For example:
half(1) returns (0, false)
half(2) returns (1, true)
*/

package main

import "fmt"

func half(a int) (int, bool) {
	quotient := a / 2
	var isEven bool
	if a%2 == 0 {
		isEven = true
	} else {
		isEven = false
	}
	return quotient, isEven
}

func main() {
	fmt.Println(half(2))
}
