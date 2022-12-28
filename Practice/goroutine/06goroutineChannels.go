package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go hello("Ajinkya", c)

	for msg := range c {

		fmt.Println(msg)
	}
}

func hello(name string, c chan string) {

	for i := 0; i < 5; i++ {

		msg := fmt.Sprintf("Good Morning %s!", name)
		c <- msg
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
