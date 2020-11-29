// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import "fmt"

func main() {

	// Declare and make a map of integer type values.
	scores := make(map[string]int)

	// Initialize some data into the map.
	scores["one"] = 1
	scores["two"] = 2
	scores["three"] = 3
	scores["four"] = 4
	scores["five"] = 5

	// Display each key/value pair.
	for k, v := range scores {
		// fmt.Println(k, v)
		fmt.Printf("Score for player: %s\t is %d\n", k, v)
	}
}
