package main

import "fmt"

func myfun(a interface{}) {

	switch a.(type) {

	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("\nType: string, Value: ", a.(string))
	case float64:
		fmt.Println("\nType: float64, Value: ", a.(float64))
	// case bool:
	// 	fmt.Println("\nType: bool, Value: ", a.(bool))
	default:
		fmt.Println("\nType not found")
	}
}

func main() {

	myfun("Radha Krishna")
	myfun(67.9)
	myfun(true)
}
