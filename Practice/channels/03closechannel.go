package main

import "fmt"

func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "Ganpati Bappa Morya"
	}
	close(mychnl)
}

func main() {

	c := make(chan string)

	go myfun(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
