package main

import "fmt"

func main() {
	colors := map[string]string{
		"blue":  "#0000FF",
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"
	// fmt.Println(colors)

	delete(colors, "white")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, value := range c {
		fmt.Println("Hex code for " + color + " is " + value)
	}
}
