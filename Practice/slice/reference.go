package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}
	nums2 := nums

	nums2[0] = 11
	nums2[1] = 22

	fmt.Println(nums)
	fmt.Println(nums2)
}
