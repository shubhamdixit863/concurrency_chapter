package main

import (
	"fmt"
	"time"
)

func Data(ch chan int) {

	fmt.Println("Sending started")
	fmt.Println("Nap time over------------")
	ch <- 1
	fmt.Println("Sending done")

}

func main() {
	ch := make(chan int)
	go Data(ch)
	time.Sleep(5 * time.Second)
	fmt.Println(<-ch)
	fmt.Println("Value received")

}

// You have to spawn 2 go routines
// is send a struct over the channel
// and print out the value of the struct in another routine
// then exit the main routine
