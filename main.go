package main

import (
	"fmt"
	"sync"
)

var value int = 10

func main() {

	var wg sync.WaitGroup
	var mutex sync.Mutex

	go func() {
		defer wg.Done()
		mutex.Lock()
		for i := 0; i < 100000; i++ {
			value++
		}
		mutex.Unlock()

	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		for i := 0; i < 100000; i++ {
			value++
		}

		mutex.Unlock()

	}()

	go func() {
		defer wg.Done()
		mutex.Lock()
		for i := 0; i < 100000; i++ {
			value--
		}

		mutex.Unlock()

	}()

	go func() {
		defer wg.Done()
		mutex.Lock()

		for i := 0; i < 100000; i++ {
			value--
		}
		mutex.Unlock()

	}()

	wg.Add(4)

	wg.Wait()
	fmt.Println(value)

}

// selection sort algorithm ---->
