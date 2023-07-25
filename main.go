package main

import "fmt"

func Data(ch chan int) {
	c := 10

	ch <- c //  <-ch
	close(ch)

}

func main() {
	ch := make(chan int) // unbuffered channel
	go Data(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
