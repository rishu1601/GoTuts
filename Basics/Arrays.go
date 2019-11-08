package main

import "fmt"

func main() {
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
