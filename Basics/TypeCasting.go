package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

// Shows how to return multiple values

//Contains information about basic types and type casting available with go
func main() {
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

	// s1, s2 := "Rishabh ", "Why "
	// fmt.Println(multipleReturns(s1, s2))
}
