//stack and everything else (heap memory)
// shared down everything stays on the stack
// each function has its own stack frame

package main

func main() {
	n := 4
	inc(&n)
	println(n)
}

func inc(x *int) {
	*x++
}
