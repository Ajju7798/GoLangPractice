package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		count("bananas")
		wg.Done()
	}()

	go func() {

		count("cherries")
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {

	for i := 0; i < 4; i++ {

		fmt.Printf("counting %s\n", thing)
		time.Sleep(time.Millisecond * 500)
	}
}
