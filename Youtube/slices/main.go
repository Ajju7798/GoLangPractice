package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices Learning")

	var fruitList = []string{"apple", "pear", "banana"}

	fmt.Printf("type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "cherry")

	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])

	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 473
	highScore[1] = 952
	highScore[2] = 549
	highScore[3] = 894
	// highScore[4] = 407

	highScore = append(highScore, 444, 555, 666)

	// fmt.Println(highScore)

	sort.Ints(highScore)

	// fmt.Println(highScore)
	// fmt.Println(sort.IntsAreSorted(highScore))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}

	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
