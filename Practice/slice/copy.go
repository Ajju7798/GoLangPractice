package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	slice2 := make([]int, len(slice))

	n := copy(slice2, slice)

	fmt.Printf("%d elements copied\n", n)

	fmt.Println("slice:", slice)
	fmt.Println("slice2:", slice2)
}
