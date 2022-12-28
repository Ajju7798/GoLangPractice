package main

import "fmt"

type User struct {
	name       string
	occupation string
	country    string
}

func main() {

	users := []User{

		{"Bunni	MacKinnon", "Product Engineer", "USA"},
		{"Leora	McLelland", "Marketing Assistant", "UK"},
		{"Dayna	Millam", "programmer", "Canada"},
		{"Marney	Medlen", "Design Engineer", "Slovakia"},
		{"Lorinda	Butcher", "Accounting Assistant IV", "USA"},
		{"Audrey	Gerardet", "programmer", "Canada"},
		{"Ulrica	Brambill", "Payment Adjustment Coordinator", "Slovakia"},
	}

	var programmers []User

	for _, user := range users {

		if isProgrammer(user) {
			programmers = append(programmers, user)
		}
	}

	fmt.Println("Programmers:")
	for _, u := range programmers {

		fmt.Println(u)
	}
}

func isProgrammer(user User) bool {

	return user.occupation == "programmer"
}
