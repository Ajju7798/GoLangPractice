/*
package main

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {


		var a Address
		fmt.Println(a)

		a1 := Address{"Ajinkya", "Pune", 123456}

		fmt.Println("Address1: ", a1)

		a2 := Address{Name: "Rutuja", city: "Mumbai",
			Pincode: 277001}

		fmt.Println("Address2: ", a2)

		a3 := Address{Name: "Ramesh"}
		fmt.Println("Address3: ", a3)

}

*/

/* package main

import "fmt"

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

func main() {
	c := Car{Name: "Tata", Model: "Altroz",
		Color: "White", WeightInKg: 1920}

	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	c.Color = "Black"

	fmt.Println("Car: ", c)
} */

/* package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	emp8 := &Employee{"Ajinkya", "Kamble", 24, 80000}

	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)
} */

/* package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	emp8 := &Employee{"Ajinkya", "Kamble", 24, 80000}

	fmt.Println("First Name: ", emp8.firstName)
	fmt.Println("Age: ", emp8.age)
} */

/* package main

import "fmt"

type User struct {
	name       string
	occupation string
	age        int
}

func main() {

	u := User{}
	u.name = "Ajinkya Kamble"
	u.occupation = "Software Developer"
	u.age = 24

	fmt.Printf("%s is %d years old and he is a %s.\n", u.name, u.age, u.occupation)
} */

// Go anonymous struct

/* package main

import "fmt"

func main() {

	u := struct {
		name       string
		occupation string
		age        int
	}{
		name:       "Ajinkya Kamble",
		occupation: "Software Developer",
		age:        24,
	}

	fmt.Printf("%s is %d years old and he is a %s.\n", u.name, u.age, u.occupation)
}


*/

// Go nested structs

/* package main

import "fmt"

type Address struct {
	city    string
	country string
	state   string
	pincode int
}

type User struct {
	name    string
	age     int
	address Address
}

func main() {

	p := User{
		name: "Ajinkya Kamble",
		age:  24,
		address: Address{
			city:    "Pune",
			state:   "Maharashtra",
			country: "India",
			pincode: 411018,
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
	fmt.Println("Country:", p.address.country)
	fmt.Println("Pincode:", p.address.pincode)
} */

// Go struct promoted fields

/* package main

import "fmt"

type Address struct {
	city    string
	country string
	state   string
	pincode int
}

type User struct {
	name string
	age  int
	Address
}

func main() {
	p := User{
		name: "Ajinkya Kamble",
		age:  24,
		Address: Address{
			city:    "Pune",
			state:   "Maharashtra",
			country: "India",
			pincode: 411018,
		},
	}

	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)
	fmt.Println("State:", p.state)
	fmt.Println("Country:", p.country)
	fmt.Println("Pincode:", p.pincode)
} */

// Go struct functions fields

/* package main

import "fmt"

type Info func(string, string, int) string

type User struct {
	name       string
	occupation string
	age        int
	info       Info
}

func main() {

	u := User{
		name:       "Ajinkya Kamble",
		occupation: "Software Developer",
		age:        24,
		info: func(name string, occupation string, age int) string {

			return fmt.Sprintf("%s is %d years old and he is a %s\n", name, age, occupation)
		},
	}

	fmt.Printf(u.info(u.name, u.occupation, u.age))
} */

// Go struct constructor

/* package main

import "fmt"

type User struct {
	name       string
	occupation string
	age        int
}

func newUser(name string, occupation string, age int) *User {

	p := User{name, occupation, age}
	return &p
}

func main() {

	u := newUser("Ramesh Raut", "painter", 44)

	fmt.Printf("%s is %d years old and he is a %s\n", u.name, u.age, u.occupation)
} */

// Go struct is a value type

/* package main

import "fmt"

type User struct {
	name       string
	occupation string
	age        int
}

func main() {

	u1 := User{"Sam Johnson", "gardener", 41}

	u2 := u1

	u2.name = "John Richard"
	u2.occupation = "driver"
	u2.age = 34

	fmt.Printf("%s is %d years old and he is a %s\n", u1.name, u1.age, u1.occupation)
	fmt.Printf("%s is %d years old and he is a %s\n", u2.name, u2.age, u2.occupation)
} */

// Comparing Go structs

/* package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {

	p1 := Point{3, 4}
	p2 := Point{3, 4}

	if p1 == p2 {

		fmt.Println("The structs are equal")
	} else {

		fmt.Println("The structs are not equal")
	}
} */

// Go create a slice of structs

/* package main

import "fmt"

type User struct {
	name       string
	occupation string
	country    string
}

func main() {

	users := []User{}
	users = append(users, User{"Adrian Abraham", "gardener", "USA"})
	users = append(users, User{"Alan Allan", "driver", "UK"})
	users = append(users, User{"Alexander Alsop", "programmer", "Canada"})
	users = append(users, User{"Andrew Anderson", "teacher", "Slovakia"})
	users = append(users, User{"Anthony Arnold", "shopkeeper", "USA"})
	users = append(users, User{"Austin Avery", "programmer", "Canada"})
	users = append(users, User{"Benjamin Bailey", "programmer", "Slovakia"})

	for _, user := range users {

		fmt.Println(user)
	}
} */

// Go filter slice of structs

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
