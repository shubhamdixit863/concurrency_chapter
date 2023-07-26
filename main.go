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
