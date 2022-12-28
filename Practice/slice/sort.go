package main

import (
	"fmt"
	"sort"
)

func main() {

	fruits := []string{"banana", "grapes", "apple", "cherry", "mango", "orange", "custord-apple"}
	nums := []int{4, 2, 1, 5, 6, 8, 0, -3}

	sort.Strings(fruits)
	sort.Ints(nums)

	fmt.Println(fruits)
	fmt.Println(nums)
}
