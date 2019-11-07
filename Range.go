//Range is used as an iterator for a lot of data structures like map,arrays,slices,strings

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range slice1 { //_ for index,as index is not required
		sum += num
	}
	fmt.Println("Sum : ", sum)

	mapVeg := map[string]string{"a": "apple", "b": "banana", "c": "cucumber"}
	for k, v := range mapVeg { //k,v stand for key,value
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "Rishabh" { //Prints the unicode
		fmt.Println(i, c)
	}
}
