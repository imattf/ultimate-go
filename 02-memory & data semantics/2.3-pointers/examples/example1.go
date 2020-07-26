// show the basic concept of pass by value.

package main

func main() {
	//declare variable of type int w/ a value of 10
	count := 10

	//display the "vaue of" and "address of" count
	println("count:\tVaule of[", count, "]\tAddr of[", &count, "]")

	//pass the "value of" count
	increment(count)

	//display the "vaue of" and "address of" count
	println("count:\tVaule of[", count, "]\tAddr of[", &count, "]")

}

func increment(inc int) {

	//increment the "value of" local inc
	inc++

	//display the "vaue of" and "address of" count
	println("inc:\tVaule of[", inc, "]\tAddr of[", &inc, "]")
}