package main

import "fmt"

func main() {
	fmt.Println("Welcome to ifElse")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"

	} else if loginCount > 10 {
		result = "watchout"
	} else {
		result = "Exactly 10 login counts"
	}

	fmt.Println(result)

}
