package main

import "fmt"

func main() {

	slice := make([]int, 3)

	fmt.Printf("slice: %v; len: %d; cap: %d \n", slice, len(slice), cap(slice))

	fmt.Println("---------------------------")

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4, 5, 6, 7, 8)

	fmt.Printf("slice: %v; len: %d; cap: %d \n", slice, len(slice), cap(slice))
}
