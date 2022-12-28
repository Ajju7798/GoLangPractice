package main

import (
	"fmt"
	"structs/model"
)

func main() {

	u := model.User{Name: "Ajinkya Kamble", Occupation: "Software Developer"}

	fmt.Printf("%s is a %s\n", u.Name, u.Occupation)
}
