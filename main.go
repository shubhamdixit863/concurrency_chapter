package main

import (
	"fmt"
	"sync"
)

func Even(ch chan int, wg *sync.WaitGroup, slc []int) {

	for _, va := range slc {
		if va%2 == 0 {
			ch <- va
		}
	}

	wg.Done()
	close(ch)
}
func Odd(ch chan int, wg *sync.WaitGroup, slc []int) {

	for _, va := range slc {
		if va%2 != 0 {
			ch <- va
		}
	}

	wg.Done()
	close(ch)

}

func main() {
	d := []int{12, 7, 1, 3}
	var even []int
	var odd []int

	var wg sync.WaitGroup
	evenChannel := make(chan int)
	oddChannel := make(chan int)
	go Even(evenChannel, &wg, d)
	go Odd(oddChannel, &wg, d)
	wg.Add(2)

	even = append(even, <-evenChannel)
	odd = append(even, <-oddChannel)

	wg.Wait()
	fmt.Println(even)
	fmt.Println(odd)

}
