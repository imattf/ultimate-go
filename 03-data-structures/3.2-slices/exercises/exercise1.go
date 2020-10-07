// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

func main() {

	// Declare a nil slice of integers.
	var xi []int

	// Append numbers to the slice.
	xi = append(xi, 1)
	xi = append(xi, 2)
	xi = append(xi, 3)

	// for i:=0; i<10; i++ {
	// 	xi = append(xi, i*10)
	// }

	// Display each value in the slice.
	for _, val := range xi {
		fmt.Println(val)
	}

	// Declare a slice of strings and populate the slice with names.
	friends := []string{"Annie", "Betty", "Charely", "Doug", "Edward"}

	// Display each index position and slice value.
	for i, v := range friends {
		fmt.Printf("index is %d and value is %v\n", i, v)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	xs := friends[1:3]

	// Display each index position and slice values for the new slice.
	for i, v := range xs {
		fmt.Printf("index is %d and value is %v\n", i, v)
	}
}
