package main

import "fmt"

func main() {

	contries := make(map[string]string)

	contries["In"] = "India"
	contries["Usa"] = "United States of America"
	contries["Uk"] = "United Kingdom"

	fmt.Println(contries)
	fmt.Printf("%q\n", contries)
}
