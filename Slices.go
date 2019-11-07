package main

import "fmt"

func main() {
	slice1 := make([]int, 3) //Make allocates and initializes an object of type slice, map, or chan
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3

	fmt.Println("Slice:", slice1)

	slice1 = append(slice1, 4)
	fmt.Println("Append:", slice1) //Appending to a slice

	slice1Copy := make([]int, len(slice1))
	copy(slice1Copy, slice1) //copy(dest,source)
	fmt.Println("Copy:", slice1Copy)

	slice2 := []int{6, 7, 8} //Initialization using braces
	fmt.Println("Initialization:", slice2)

	slicedPart := slice1[:2]
	fmt.Println("Sliced:", slicedPart)

	twoDimSlice := make([][]int, 3) //no. of rows = 3
	for i := 0; i < 3; i++ {
		innerLen := i + 1 // No. of cols = 1,2,3
		twoDimSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDimSlice[i][j] = i + j
		}
	}
	fmt.Println("2 D Slice: ", twoDimSlice)

}
