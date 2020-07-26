//literal structs
//type pollution...
// you don't have to name something
// if it's only gonna be used in one place
// just define it on the fly

package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	//Declare a variable of an anonymous type
	// and set it to its zero value
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	//Display the value; special display patterns for structs
	fmt.Printf("%v\n", e1)  //just values
	fmt.Printf("%+v\n", e1) //labels & values
	fmt.Printf("%#v\n", e1) //function, label and values

	//Declare a variable of anonymous type and
	// init using a struct literal
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 1,
		pi:      3.141593,
	}

	//Display the field value
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

}
