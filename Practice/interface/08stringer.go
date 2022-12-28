package main

import (
	"fmt"
)

type User struct {
	Name       string
	Occupation string
}

func (u User) String() string {

	return fmt.Sprintf("%s is a(n) %s", u.Name, u.Occupation)
}

func main() {

	u1 := User{"Ajinkya Kamble", "software Developer"}
	u2 := User{"Pratik Dahiya", "site manager"}

	fmt.Println(u1)
	fmt.Println(u2)
}
