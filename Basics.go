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

//main will always run-> similar to c/c++/java
func main() {
	// The package functions are called with a Capitalized func name
	fmt.Println("Hello World!Here at go now")
	fmt.Println("Random Value : ", rand.Intn(101))
	// types()
	// pointers()
	// forLoopInGo()
	// switchInGo()
	// ifElseInGo()
	// arraysInGo()
}

func add(x, y float64) float64 {
	return x + y
}

// Shows how to return multiple values
func multipleReturns(a, b string) (string, string) {
	return a + "Hi", b + "There"
}

//Contains information about basic types and type casting available with go
func types() {
	//Types cannot be dynamic in go, types once compiled are the same throughout the lifecycle
	var num1 float64 = 8.12
	var num2 float64 = 9.32

	var b int = 10
	var c float64 = float64(b) //Typecasting

	x := b // x will be of type int
	//:= is used instead of = for the variable assignment

	fmt.Println(b, c, x)

	//Shorthand Syntax
	var num3, num4 float64 = 9.213, 9.312

	//Inside function,no need to specify types
	num5, num6 := 321.321, 123.312

	fmt.Println(add(num1, num2))
	fmt.Println(add(num3, num4))
	fmt.Println(add(num5, num6))

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

//switch case in go
func switchInGo() {
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

//arrays in go
func arraysInGo() {
	var a [5]int
	fmt.Println("initial:", a) //0 indexed arrays

	a[4] = 100 //setting the last value

	fmt.Println("Set: ", a) //[0 0 0 0 100]

	fmt.Println("Length: ", len(a))

	var b [2][3]int //2-d array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = i + j
		}
	}
	fmt.Println("2-d array: ", b)
}
