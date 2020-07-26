//show the basic concept of using a pointer
// to share data.

package main

import "fmt"

//user type (Class)
type user struct {
	name   string
	email  string
	logins int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u := user{
		name:  "bob",
		email: "bob@aol.com",
	}

	//pass the address of the u value
	display(&u)

	//pass the "value of" count
	increment(&u.logins)

	//pass the address of the u value
	display(&u)

}

// increment declares logins as a pointer variable whose value is
// always an address and points to values of type int.
func increment(inc *int) {

	//increment the "value of" local inc
	*inc++

	//display the "vaue of" and "address of" count
	println("inc:\tVaule of[", inc, "]\tAddr of[", &inc, "]\tValue Points to[", *inc, "]")
}

// display declares u as user pointer variable whose value is always an address
// and points to values of type user.
func display(u *user) {
	fmt.Printf("%p\t%+v\n", u, *u)
	fmt.Printf("Name: %q Email: %q Logins: %d\n\n", u.name, u.email, u.logins)

}
