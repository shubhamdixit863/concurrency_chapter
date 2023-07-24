package main

import (
	"fmt"
	"sync"
)

func Data(w *sync.WaitGroup) {
	fmt.Println("jejeje")
	for i := 0; i < 10; i++ {
		fmt.Println("Running from go routine", i)
	}

	w.Done()
}

func main() {

	var wg sync.WaitGroup

	go Data(&wg)

	wg.Add(1)
	fmt.Println("Hello world")

	//time.Sleep(1 * time.Second)  // wrong way

	//fmt.Scanln()   // wrong way

	// Correct way is
	// channels
	// waitGroups
	wg.Wait()
	fmt.Println("main routine exited")

}
