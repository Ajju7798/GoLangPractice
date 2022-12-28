package main

import "fmt"

func uniq(slice []int) []int {

	uslice := []int{}
	seen := make(map[int]bool)

	for _, val := range slice {

		if _, in := seen[val]; !in {

			seen[val] = true
			uslice = append(uslice, val)
		}
	}

	return uslice
}

func main() {

	slice := []int{1, 2, 2, 3, 4, 4, 5, 6, 7, 8, 8, 8, 9, 9}
	uslice := uniq(slice)

	fmt.Printf("Original slice: %v\n", slice)
	fmt.Printf("Unique slice: %v\n", uslice)
}
