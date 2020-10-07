// Sample program to show how to declare and iterate over
// arrays of different types.

package main

import "fmt"

func main() {
	// Declare an array of five strings that is
	// initialized to its zero value.

	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// Iterate over the array of strings.
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

	// Declare an array of 4 integrers that is
	// initialized with some non-zero values.

	numbers := [4]int{10, 20, 30, 40}

	// Iterate over the array of integers.
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}

}
