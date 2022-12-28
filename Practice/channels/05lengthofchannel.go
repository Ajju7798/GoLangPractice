package main

import "fmt"

func main() {

	mychnl := make(chan string, 5)
	mychnl <- "one"
	mychnl <- "two"
	mychnl <- "three"
	mychnl <- "four"

	fmt.Println("Length of the channel is: ", len(mychnl))
	fmt.Println("Capacity of the channel is: ", cap(mychnl))
}
