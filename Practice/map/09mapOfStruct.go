package main

import "fmt"

type User struct {
	name       string
	occupation string
}

func main() {

	u1 := User{
		name:       "Frank N. Stein",
		occupation: "photographer",
	}

	u2 := User{
		name:       "Joe V. Awl",
		occupation: "site manager",
	}

	u3 := User{
		name:       "Col Fays",
		occupation: "actor",
	}

	users := map[int]User{
		1: u1,
		2: u2,
		3: u3,
	}

	fmt.Println(users)
}
