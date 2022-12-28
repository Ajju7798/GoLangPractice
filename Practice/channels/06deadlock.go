package main

import "fmt"

func main() {

	c := make(chan string)

	c <- "lenovo"
	msg := <-c

	fmt.Println(msg)
}
