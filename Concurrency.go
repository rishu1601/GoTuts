package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //Notifies that work has been done
}

// func main() {
// 	//goroutine is like a lightweight thread
// 	//to make a function concurrent,just add go infront of the function call

// 	go say("Hi")
// 	say("Why")

// 	/*
// 		goroutine is non blocking, so if program execution finishes before goroutine
// 		we cannot have the functioncall completed
// 	*/
// 	//Try this

// 	// go say("hi")
// 	// go say("there")

// 	//Try this

// 	// say("hi")
// 	// go say("why")
// }

//Goroutine synchronization

var wg sync.WaitGroup

func main() {
	wg.Add(1) //Adds to waitgroup
	go say("hi")
	wg.Add(1)
	go say("no")
	wg.Wait() //Asks for program to wait
}
