package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()
	r := recover() // listens to panic and lets the program run smooth
	if r != nil {
		fmt.Println("Recovered here ", r)
	}
}
func say(s string) {
	defer cleanup()
	for i := 10; i < 20; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 15 {
			panic("Its a 15") //Everything crashes after this,i.e, the program stops executing
		}
	}
}

func main() {
	wg.Add(1) //Adds to waitgroup
	go say("hi")
	wg.Add(1)
	go say("no")
	wg.Wait() //Asks for program to wait
}
