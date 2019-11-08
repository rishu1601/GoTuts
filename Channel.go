package main

import "fmt"

func foo(c chan int, someval int) {
	c <- someval * 5 //returns 5 * someval to the channel c
}

func main() {
	c := make(chan int)
	//send and receive of a channel is blocking
	//So we dont need to synchronise these goroutines
	// go foo(c, 3)
	// go foo(c, 5)

	// v1, v2 := <-c, <-c
	// fmt.Println(v1, v2) //Prints 25,15 ... LIFO order

	//Buffering and iterating over channels
	for i := 0; i < 10; i++ {
		go foo(c, i)
	}
	for item := range c {
		fmt.Println(item)
		/*
			Gives the following error:
			fatal error: all goroutines are asleep - deadlock!
				goroutine 1 [chan receive]:

			//This is because go routines dont comeback before the program ends
			To get rid of this error we need to synchronise the goroutines
		*/
	}
}
