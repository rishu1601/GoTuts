package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	//goroutine is like a lightweight thread
	//to make a function concurrent,just add go infront of the function call

	go say("Hi")
	say("Why")

	/*
		goroutine is non blocking, so if program execution finishes before goroutine
		we cannot have the functioncall completed
	*/
	//Try this

	// go say("hi")
	// go say("there")

	//Try this

	// say("hi")
	// go say("why")
}
