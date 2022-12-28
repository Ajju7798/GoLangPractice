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

	new_map := map1

	new_map[106] = "Parrot"
	new_map[108] = "Pig"

	fmt.Println("New map: ", new_map)
	fmt.Println("\nModification done in old map:\n", map1)
}
