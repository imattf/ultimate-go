// Sample program to show how only types that can have
// equality defined on them can be a map key.

package main

import "fmt"

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

func main() {

	// Declare and initialize the map with values.
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Iterate over the map.
	for key, value := range users {
		fmt.Println(key, value)
	}

	fmt.Println()

	// Iterarte over the map printinng just the keys.
	// Notice random results
	for key := range users {
		fmt.Println(key)
	}

}
