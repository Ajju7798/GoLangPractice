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

	delete(map1, 100)
	delete(map1, 103)

	fmt.Println("Map after deletion: ", map1)
}
