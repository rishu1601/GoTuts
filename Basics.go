package main

// Everything starts with package main

//For multiple package imports, we use this syntax
import (
	"fmt"
	"math/rand"
)

//In go lang we follow /subpackage notation for subpackage imports
// Every import has to be with the double quotes
// "fmt" is for format
func multipleReturns(a, b string) (string, string) {
	return a + "Hi", b + "There"
}

//main will always run-> similar to c/c++/java
func main() {
	// The package functions are called with a Capitalized func name
	fmt.Println("Hello World!Here at go now")
	fmt.Println("Random Value : ", rand.Intn(101))

	s1, s2 := "Rishabh ", "Why "
	fmt.Println(multipleReturns(s1, s2))

}

//Pointers in go
func pointers() {
	x := 10
	a := &x         //a stores the memory address of x
	fmt.Println(a)  //address of x
	fmt.Println(*a) //value of x
	*a = 5
	fmt.Println(x) //value of x is changed
}

//for loop in go
func forLoopInGo() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//Print odd numbers
	fmt.Println("Odd values less than 5")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func ifElseInGo() {
	//NOTE=> There is no ternary if else in go
	num := 9
	if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Positive")
	}
	//No parentheses are required but braces are mandatory
}
