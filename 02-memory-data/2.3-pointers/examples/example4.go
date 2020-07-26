package main
/* escape analysis...

frame / stack mental model
stacks are self-cleaning

value syamntics provides...
- isolation model
- mutation to happen in safeway
- data locality
- but the in-efficiancy of it all
we leverage pointer symantics...


complier does static code analysis
  one these is escape analyis
compiler determines where value should be created in memory (stack or heap)
  when value is created on the heap, we call that an escape
  when function exits, its frame on memory on the stack will be cleaned,
	if there is a pointer reference still needed by in the program elsewhere,
	compiler will create the value on the heap,
	so program needs to pass the address value back up the stack
anything on the heap is managed with the garbage collector

need to imbed //go:noline for above command to work??
go build -gcflags -m=2 will not build, but will create escape analysis report.
  needs package to work, so I did..
go run -gcflags "-m" example4.go to make this work

*/

/* - guidelines
always use value construction, not address construction
*/

// see https://play.golang.org/p/n9HijcdZ3pT

// user represents a user in the system
type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

// create a user value pass a copy back to caller
//  value symantics - getting a copy of your value
//  using factory functions
//go:noinline
func createUserV1() user {
	u := user{
		name:  "bob",
		email: "bob@aol.com",
	}
	println("V1", &u)
	return u
}

// create a user value and shares the value back to caller
//  pointer symantics - getting shared access to your value
//  using factory functions
//go:noinline
func createUserV2() *user {
	u := user{
		name:  "bev",
		email: "bev@aol.com",
	}
	println("V2", &u)
	return &u
}
