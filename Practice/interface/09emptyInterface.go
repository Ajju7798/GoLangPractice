package main

import (
	"fmt"
)

type Body struct {
	Msg interface{}
}

func main() {

	b := Body{"Hello there"}
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)

	b.Msg = 9
	fmt.Printf("%#v %T\n", b.Msg, b.Msg)
}
