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

	value_1 := map1[100]
	value_2 := map1[103]
	fmt.Println("Value of key[100]: ", value_1)
	fmt.Println("Value of key[103]: ", value_2)
}
