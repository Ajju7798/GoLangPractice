package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array learning")

	var fruitList [4]string
	fruitList[0] = "apple "
	fruitList[1] = "pear"
	// fruitList[0] = "apple"
	fruitList[3] = "banana"

	fmt.Println("Fruit list is : ", fruitList)
	fmt.Println("length", len(fruitList))

	var vegList = [5]string{"potato", "beans", "pumkin"}

	fmt.Println("vegy list", vegList)
	fmt.Println("vegy list is:", len(vegList))
}
