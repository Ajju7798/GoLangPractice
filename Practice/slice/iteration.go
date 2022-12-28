package main

import "fmt"

func main() {

	fruits := []string{"banana", "grapes", "apple", "cherry", "mango", "orange", "custord apple"}

	for i := 0; i < len(fruits); i++ {

		fmt.Println(fruits[i])
	}
}
