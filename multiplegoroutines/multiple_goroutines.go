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

/**
package main

import (
	"fmt"
	"sync"
)

func Data(ch chan int) {
	fmt.Println("Sending started")
	fmt.Println("Nap time over------------")
	ch <- 1
	fmt.Println("Sending done")

}

type Person struct {
	age     int
	isVoter bool
}

func SendDer(ch chan<- *Person, wg *sync.WaitGroup) { // send only
	g := &Person{
		age:     10,
		isVoter: true,
	}
	ch <- g

	g.isVoter = true
	g.age = 100
	defer wg.Done()
	close(ch)

}

func Receiver(ch <-chan *Person, wg *sync.WaitGroup) { //receive only channel
	//v := <-ch
	fmt.Println("Receiver Go routine---", <-ch)
	defer wg.Done()
}

func main() {
	ch := make(chan *Person)

	var wg sync.WaitGroup
	go Receiver(ch, &wg)
	go SendDer(ch, &wg)

	wg.Add(2)
	//time.Sleep(2 * time.Second)

	wg.Wait()

}

// You have to spawn 2 go routines
// is send a struct over the channel
// and print out the value of the struct in another routine
// then exit the main routine

// We can do buubble sort using channels ,select statement and go routines->

func Data(ch chan int) {

	fmt.Println("Started----")
	ch <- 1
	fmt.Println("Sent the value----")
	fmt.Println("Finished")
	close(ch)
}

func main() {

	ch := make(chan int)
	go Data(ch)
	time.Sleep(5 * time.Second)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"

		close(c1)

	}()

	go func() {
		time.Sleep(2 * time.Second)

		c2 <- "two"

	}()
loop:
	for {
		select {
		case msg := <-c1:
			if msg == "" {
				break loop
			}

			fmt.Println("Value received from channe1", msg)
		case <-c2:
			fmt.Println("Value recieved from channel2")

		}

	}

	//fmt.Println("Main finished")


*/
