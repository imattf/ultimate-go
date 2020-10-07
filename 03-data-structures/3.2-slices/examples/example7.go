// Sample program to show how to declare and use variadic functions.
package main

import (
	"fmt"
)

// user is a struct type that declares user information.
type user struct {
	id int
	name string
}

func main() {
	// Declare and initialize a value of type user
	u1 := user{id:1432, name:"Betty",}
	u2 := user{id:5678, name:"Bertie",}

	// Display both values.
	display(u1,u2)

	// Create a slice of user values.
	u3 := []user{
		{123, "Bill"},
		{456, "Joey"},
	}

	// Display all three values from the slice.
	display(u3...)

	change(u3...)
	fmt.Println("**********************")
	for _, u := range u3 {
		fmt.Printf("%+v\n", u)
	}

}

// Display can accept and display multiple values of user types
func display(users ...user) {
	fmt.Println("**********************")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

}

// Change can show how the backing array is shared
func change(users ...user) {
	users[1] = user{99, "Same Backing Array"}
}