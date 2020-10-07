// Sample program to show the components of a slice. It has a
// length, capacity and the underlying array.

package main

import "fmt"

func main() {

	// Create a slice of length 5 and capacity of 8
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	inspectSlice(fruits)

}
// inspectSlice exposes the slice header for review
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
}