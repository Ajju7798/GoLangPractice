package main

import "fmt"

func main() {

	var My_map = make(map[float64]string)
	fmt.Println(My_map)

	My_map[1.1] = "Ajinkya"
	My_map[1.2] = "Ramesh"
	fmt.Println(My_map)
}
