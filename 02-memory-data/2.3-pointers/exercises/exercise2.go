// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.

// Declare a type named user.
type user struct {
	name  string
	email string
	admin bool
}

// Create a function that changes the value of one of the user fields.
func funcName(u *user, v bool) {

	// Use the pointer to change the value that the
	// pointer points to.
	u.admin = v

}

func main() {

	// Create a variable of type user and initialize each field.
	u1 := user{
		name:  "bob",
		email: "bob@aol.com",
		admin: false,
	}

	// Display the value of the variable.
	println(u1.name, u1.email, u1.admin)

	// Share the variable with the function you declared above.
	funcName(&u1, true)

	// Display the value of the variable.
	println(u1.name, u1.email, u1.admin)
}
