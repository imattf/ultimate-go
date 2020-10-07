// Sample program to show how arrays of different sizes are
// not of the same type.

package main

import "fmt"

func main() {

	// Declare an array of 5 integrers that is
	// initialized with zero value.
	var five [5]int

	// Declare an array of 4 integrers that is
	// initialized with non-zero value.
	four := [4]int{10, 20, 30, 40}

	// Assign one array to the other
	five = four
	// Cannot use four (type [4]int) as type [5]int in assignment
	fmt.Println(four)
	fmt.Println(five)

}
