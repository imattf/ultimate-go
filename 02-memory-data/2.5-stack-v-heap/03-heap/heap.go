//stack and everything else (heap memory)
// shared up, then pointer written to heap memory

package main

func main() {
	n := answer()
	println(*n / 2)
}

func answer() *int {
	x := 42
	return &x
}
