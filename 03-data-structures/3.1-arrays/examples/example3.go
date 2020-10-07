// Sample program to show how the behavior of the for range and
// how memory for an array is contiguous.

package main

import "fmt"

func main() {

	// Declare an array of 5 strings with values.
	numbers := [5]string{"10", "20", "30", "40", "50"}

	// Iterate over the array displaying
	// the value and address of each element.
	for i, v := range numbers {
		fmt.Printf("value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &numbers[i])
	}

}
