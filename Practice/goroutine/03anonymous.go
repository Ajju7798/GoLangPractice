package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome!!! It's a Main function")

	go func() {

		fmt.Println("Ganpati Bappa Morya!!!")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! to Main function")
}
