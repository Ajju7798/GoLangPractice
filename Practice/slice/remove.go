package main

import (
	"fmt"
)

func main() {

	fruits := []string{"banana", "grapes", "apple", "cherry", "mango", "orange", "custord apple"}
	fmt.Println(fruits)

	fruits = append(fruits[1:2], fruits[2:]...)
	fmt.Println(fruits)

	fruits = append(fruits[:2], fruits[4:]...)
	fmt.Println(fruits)
}
