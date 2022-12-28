package main

import "fmt"

func main() {

	s1 := [][]string{
		{"soil", "water"},
		{"white", "black"},
		{"heaven", "hell"},
	}

	fmt.Printf("slice: %v; len: %d; cap: %d \n", s1, len(s1), cap(s1))
}
