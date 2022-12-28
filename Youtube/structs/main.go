package main

import "fmt"

func main() {
	fmt.Println("Welcome to Structs")

	// No inheritance in golang; No super or parent

	ajinkya := User{"ajinkya", "aju@gmail.com", true, 24}

	fmt.Println(ajinkya)
	fmt.Printf("ajinkya's details are: %+v", ajinkya)
	fmt.Printf("ajinkya's age is : %v \n email is : %v", ajinkya.Age, ajinkya.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
