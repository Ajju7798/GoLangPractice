package main

import (
	"fmt"
	"strings"
)

func main() {

	words := []string{}
	words = append(words, "ganpati")
	words = append(words, "bappa")
	words = append(words, "morya!!")

	res := strings.Join(words, " ")
	fmt.Println(res)
}
