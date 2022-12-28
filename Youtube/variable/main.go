package main

import "fmt"

func main() {
	// var username string = "ajinkya"
	// var isLoggedIn bool = false
	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("variables if of type: %T \n", smallVal)

	var smallFloat float64 = 256.437647575837874637
	fmt.Println(smallFloat)
	fmt.Printf("variables if of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variables if of type: %T \n", anotherVariable)

	// implicit type

	var websiter = "abc.com"

	fmt.Println(websiter)

	// no var style

	noOfUser := 300000

	fmt.Println(noOfUser)
}
