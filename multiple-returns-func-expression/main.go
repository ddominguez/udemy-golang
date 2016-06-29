/*
Write a func expression which takes an integer. The func will have two returns.
The first return should be the argument divided by 2.
The second return should be a bool that let’s us know whether or not the argument was even.
For example:
half(1) returns (0, false)
half(2) returns (1, true)
*/

package main

import "fmt"

func main() {
	half := func(a int) (int, bool) {
		return a / 2, a%2 == 0
	}
	fmt.Println(half(2))
}
