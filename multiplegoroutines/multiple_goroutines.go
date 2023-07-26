package multiplegoroutines

import (
	"fmt"
	"sync"
	"time"
)

func Data(w *sync.WaitGroup) {
	fmt.Println("jejeje")
	for i := 0; i < 10; i++ {
		fmt.Println("Running from go routine", i)
	}

	w.Done()
}

func Multiple(wg *sync.WaitGroup) {

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			time.Sleep(2 * time.Second)
			fmt.Println("The go routine", i)
			wg.Done()
		}(i)

	}

}

func main() {
	var wg sync.WaitGroup
	//go Data(&wg)

	Multiple(&wg)

	//wg.Add(1)
	fmt.Println("Hello world")

	//time.Sleep(1 * time.Second)  // wrong way

	//fmt.Scanln()   // wrong way

	// Correct way is
	// channels
	// waitGroups
	wg.Wait()
	fmt.Println("main routine exited")

	ch := make(chan int) // unbuffered channel

	ch <- 1

	value, ok := <-ch
	fmt.Println(value, ok)
	close(ch)

	c, ok := <-ch
	fmt.Println(c, ok)

}
