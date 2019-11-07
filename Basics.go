package main

// Everything starts with package main

//For multiple package imports, we use this syntax
import (
	"fmt"
	"math"
	"math/rand"
)

//In go lang we follow /subpackage notation for subpackage imports
// Every import has to be with the double quotes
// "fmt" is for format

//main will always run-> similar to c/c++/java
func main() {
	// The package functions are called with a Capitalized func name
	fmt.Println("Hello World!Here at go now")
	fmt.Println(math.Abs(-10))
	fmt.Println("The square root of 4.9 is : ", math.Sqrt(4.9))
	fmt.Println("Random Value : ", rand.Intn(101))
	// types()
	pointers()
}

func add(x, y float64) float64 {
	return x + y
}

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
