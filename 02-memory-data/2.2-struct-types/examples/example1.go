//structs
package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	//Declare a variable of type example
	// and with zero valye
	var e1 example

	//Display the value; special display patterns for structs
	fmt.Printf("%v\n", e1)  //just values
	fmt.Printf("%+v\n", e1) //labels & values
	fmt.Printf("%#v\n", e1) //function, label and values

	//Declare a variable of type example
	// with a struct literal
	e2 := example{
		flag:    true,
		counter: 1,
		pi:      3.141593,
	}

	//Display the field value
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

}
