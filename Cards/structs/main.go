package main

import "fmt"

type contact struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	// ajinkya := person{
	// 	firstName: "ajinkya",
	// 	lastName:  "kamble",
	// 	contact: contact{
	// 		email: "ajju@gmail.com",
	// 		zip:   411018,
	// 	},
	// }
	// ajinkya.updateName("ajju")
	// ajinkya.print()
	mySlice := []string{"hi", "there", "how", "are", "You"}
	updateSlice(mySlice)
	fmt.Printf("mySlice: %v\n", mySlice)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(name string) {
	(*pointerToPerson).firstName = name
}

func updateSlice(s []string) {
	s[0] = "bye"
}
