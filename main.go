package main

import (
	"fmt"
	"sync"
)

func main() {
	var value int
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			value++
		}

	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			value--
		}

	}()

	wg.Add(2)
	fmt.Println(value)
	wg.Wait()

}

// selection sort algorithm ---->
