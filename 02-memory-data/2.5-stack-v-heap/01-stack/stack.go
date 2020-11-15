//stack and everything else (heap memory)
// shared down everything stays on the stack
// each function has its own stack frame

package main

func main() {
	n := 4
	n2 := square(n)
	println(n2)
}

func square(x int) int {
	return x * x
}
