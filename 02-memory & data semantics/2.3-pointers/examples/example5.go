// show how stacks grow/change
//  share it down the call stack 10 time

/*
stacks are contiguous block of memory
engineering, nothing is free
go routines get their own stack, 
  if go routines have to share values that value has to be on heap
*/
package main

// Number of elements to grow each stack frame
const size = 1024

// main entry point for the application
func main() {
  s := "HELLO"
  stackCopy(&s, 0, [size]int{})

}

// stackCopy recursively runs increasing the size of the stack
//go:noline
func stackCopy(s *string, c int, a [size]int) {
  println(c, s, *s)

  c++
  if c == 10 {
    return
  }

  stackCopy(s, c, a)

}