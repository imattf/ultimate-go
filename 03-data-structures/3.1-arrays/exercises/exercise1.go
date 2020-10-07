// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

func main() {
	// Declare an array of 5 strings set to its zero value.
	var thing1 [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	thing2 := [5]string{"10", "20", "30", "40", "50"}

	// Assign the populated array to the array of zero values.
	thing1 = thing2

	// Iterate over the first array declared.
	// Display the string value and address of each element.
	fmt.Println(thing1)

	for i, v := range thing1 {
		fmt.Println(v, &thing1[i])
	}

}
