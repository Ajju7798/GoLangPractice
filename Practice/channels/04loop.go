package main

import "fmt"

func main() {

	mychnl := make(chan string)

	go func() {
		mychnl <- "one"
		mychnl <- "two"
		mychnl <- "three"
		mychnl <- "four"
		close(mychnl)
	}()

	for res := range mychnl {
		fmt.Println(res)
	}
}
