package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)
}

func (englishBot) getGreetings() string {
	return "Hello There!!"
}

func (spanishBot) getGreetings() string {
	return "Hola!!!"
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

// func printGreetings(sb spanishBot) {
// 	fmt.Println(sb.getGreetings())
// }
