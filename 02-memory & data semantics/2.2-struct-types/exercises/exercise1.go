// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u := user{
		name:  "bob",
		email: "bob@aol.com",
		age:   42,
	}

	// Display the field values.
	fmt.Println("Name", u.name)
	fmt.Println("email", u.email)
	fmt.Println("age", u.age)

	// Declare a variable using an anonymous struct.
	u2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "bev",
		email: "bev@aol.com",
		age:   63,
	}

	// Display the field values.
	fmt.Println("Name", u2.name)
	fmt.Println("email", u2.email)
	fmt.Println("age", u2.age)
}
