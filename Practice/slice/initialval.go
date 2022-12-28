package main

import "fmt"

func main() {

	var nums []int

	if nums == nil {
		fmt.Printf("slice is nil\n")
	}

	fmt.Printf("slice: %v; len: %d; cap: %d \n", nums, len(nums), cap(nums))

	fmt.Println("---------------------------")

	var nums2 = make([]int, 5)
	fmt.Printf("slice: %v; len: %d; cap: %d \n", nums2, len(nums2), cap(nums2))
}
