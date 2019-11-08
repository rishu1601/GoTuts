package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	//What would happen if control never reaches here? -> infinite wait

}

/*Defer

Defer statement gets evaluated whenever it is hit
defer function or method’s call arguments evaluate instantly, but they execute until the nearby functions returns
Defer runs only when its surrounding functions have evaluated,i.e defer makes sure that something runs regardless of flow control
*/

// func foo() {
// 	defer fmt.Println("Done!")         //This runs only when print in the next line is done
// 	defer fmt.Println("Are you done?") //Defer works in LIFO order
// 	fmt.Println("do something")
// }
func main() {
	wg.Add(1) //Adds to waitgroup
	go say("hi")
	wg.Add(1)
	go say("no")
	wg.Wait() //Asks for program to wait
	// foo()
}
