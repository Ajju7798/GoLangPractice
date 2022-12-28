package main

import "fmt"

func main() {

	map1 := map[int]string{
		100: "Dog",
		101: "Cat",
		102: "Cow",
		103: "Bird",
		104: "Rabbit",
	}

	fmt.Println("Original map: ", map1)

	pet_name, ok := map1[100]
	fmt.Println("\nKey present or not:", ok)
	fmt.Println("Value:", pet_name)

	_, ok1 := map1[102]
	fmt.Println("\nKey present or not:", ok1)
}
