package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someval int) {
	defer wg.Done()
	c <- someval * 5 //returns 5 * someval to the channel c
}

func main() {
	c := make(chan int, 10) //10 indicates that we are buffering for 10 items
	for i := 0; i < 10; i++ {
		//Synchronise the waitgroup
		wg.Add(1)
		go foo(c, i)
	}
	wg.Wait()
	//Wait for all goroutines to finish before closing the channels
	close(c)
	for item := range c {
		fmt.Println(item)
	}
}
