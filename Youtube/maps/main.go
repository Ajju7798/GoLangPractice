package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps")

	languages := make(map[string]string)

	languages["Js"] = "Javascript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for ", languages["Js"])
	delete(languages, "Rb")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang

	for _, value := range languages {
		fmt.Printf("For  value is %v\n", value)
	}
}
