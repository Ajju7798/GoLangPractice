package main

import "fmt"

func main() {

	var map_1 map[int]int

	if map_1 == nil {

		fmt.Println("True")
	} else {

		fmt.Println("False")
	}

	map_2 := map[int]string{

		101: "Dog",
		102: "Cat",
		103: "Cow",
		104: "Bird",
		105: "Rabbit",
	}

	fmt.Println("Map-2: ", map_2)
}
