package main

import "fmt"

func main() {

	m_a_p := map[int]string{

		100: "Dog",
		101: "Cat",
		102: "Cow",
		103: "Bird",
		104: "Rabbit",
	}

	for id, pet := range m_a_p {

		fmt.Println(id, pet)
	}
}
