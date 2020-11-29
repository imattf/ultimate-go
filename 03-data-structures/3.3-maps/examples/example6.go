// Sample program to show that you cannot take the address
// of an element in a map.

package main

// player represents someone playing our game
type player struct {
	name  string
	score int
}

func main() {

	// Declare a map with initial values using a map literal
	players := map[string]player{
		"anna":  {"Anna", 21},
		"jacob": {"Jacob", 42},
	}
	// Try to take the address of a map element; it fails
	anna := &players["anna"]
	anna.score++

	// Instead take the element, modify it, and put it back.
	player := players["anna"]
	player.score++
	players["anna"] = player
}
