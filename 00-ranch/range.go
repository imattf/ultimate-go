//home on the range

package main

import "fmt"

func main() {
	//make a map
	meals := make(map[string]string)

	//add map values
	meals["breakfast"] = "eggs"
	meals["lunch"] = "samich"
	meals["dinner"] = "pizza"

	//range for the map
	for i, v := range meals {
		fmt.Printf("%+v and %+v \n", i, v)

	}

	//make a slice of string
	greetings := []string{"ciao", "bonjour", "hi"}

	//range for the slice
	for i, v := range greetings {
		fmt.Printf("%+v and %+v \n", i, v)

	}

}
