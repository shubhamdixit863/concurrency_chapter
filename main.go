package main

import (
	"fmt"
)

func BuubleSort(slc []int) {

	for i := 0; i < len(slc); i++ {
		// inneer loop

		for j := 0; j < len(slc)-1; j++ {
			if slc[j] > slc[j+1] {
				// swap here
				slc[j], slc[j+1] = slc[j+1], slc[j]
			}
		}

	}
}

func BuubleSortConcurrent(slc []int, ch chan int, exit chan struct{}) {

	for i := 0; i < len(slc); i++ {

		ch <- i
	}

	exit <- struct{}{}
}

func InnerLoop(slc []int) {
	for j := 0; j < len(slc)-1; j++ {
		if slc[j] > slc[j+1] {
			// swap here
			slc[j], slc[j+1] = slc[j+1], slc[j]
		}
	}

}
func main() {

	c := []int{34, 1, 23, 234, 0, 2, 99}
	exit := make(chan struct{})
	//BuubleSort(c)

	buubleSortIndex := make(chan int)

	go BuubleSortConcurrent(c, buubleSortIndex, exit)

loop:
	for {
		select {
		case <-buubleSortIndex:
			InnerLoop(c)

		case <-exit:
			break loop

		}

	}

	fmt.Println("Sorted array", c)
}

// selection sort algorithm ---->
