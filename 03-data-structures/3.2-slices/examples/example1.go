//Sample program to show how the capacity of the slice
// is not available for use

package main

import "fmt"

func main() {

	// Create a slice of length 5
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// You can't access an index of a slice beyond it's length
	fruits[5] = "Runtime error"
	// -> panic: runtime error: index out of range [5] with length 5

	fmt.Println(fruits)
}