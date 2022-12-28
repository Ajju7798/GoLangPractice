package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is : ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("ptr value is : ", ptr)
	fmt.Println("ptr actual value is : ", *ptr)

	*ptr = *ptr * 2

	fmt.Println("new value is:", myNumber)

}
