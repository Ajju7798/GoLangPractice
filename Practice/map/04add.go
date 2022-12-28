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

	fmt.Println("Original map: ", m_a_p)

	m_a_p[105] = "Parrot"
	m_a_p[106] = "Crow"
	fmt.Println("Map after adding new key-value pair:\n", m_a_p)

	m_a_p[101] = "PIG"
	m_a_p[103] = "DONKEY"
	fmt.Println("\nMap after updating values of the map:\n", m_a_p)
}
