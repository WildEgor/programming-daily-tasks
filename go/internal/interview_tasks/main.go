package main

import (
	"fmt"
	"sync"
)

type in struct {
	value    int32
	oddChan  chan int32
	evenChan chan int32
}

var serverChan chan in = make(chan in, 9)

/*
 * Complete the 'Server' function below and missing types and global variables.
 *
 * The function is void.
 */

func Server(wg *sync.WaitGroup) {
	defer wg.Done()

	value := <-serverChan
	if value.value%2 == 0 {
		value.evenChan <- value.value
	} else {
		value.oddChan <- value.value
	}

}

func main() {
	var arr []int32 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}

	oddChan := make(chan int32)
	evenChan := make(chan int32)
	odds, evens := []int32{}, []int32{}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for idx := 0; idx < len(arr); idx++ {
			i := idx
			serverChan <- in{arr[i], oddChan, evenChan}
		}

		//close(oddChan)
		//close(evenChan)
	}()

	wg.Add(1)
	go Server(&wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for newOdd := range oddChan {
			odds = append(odds, newOdd)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for newEven := range evenChan {
			evens = append(evens, newEven)
		}
	}()

	wg.Wait()

	for _, resultItem := range odds {
		fmt.Printf("%d \n", resultItem)
	}

	for _, resultItem := range evens {
		fmt.Printf("%d \n", resultItem)
	}
}
