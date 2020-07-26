//type-conversion
package main

import "fmt"

func main() {

	//both takes 8 bytes of memory...

	//this can rep values to 63 used for data, 1 bit use for plus/minus
	var signedInt int
	//all bits used for values
	var unsignedInt uint

	fmt.Println("h...", signedInt, unsignedInt)

	signedInt = 1
	unsignedInt = 2
	signedInt = int(unsignedInt)  //explict conversion
	unsignedInt = uint(signedInt) //explict conversion

	fmt.Println("h...", signedInt, unsignedInt)
	fmt.Printf("%+v and %+v \n", signedInt, unsignedInt)

}
