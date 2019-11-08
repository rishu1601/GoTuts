package main

import "fmt"

func main() {

	m := make(map[string]int) //map["key-type"]"val-type"
	m["a"] = 0
	m["b"] = 1
	fmt.Println("Map:", m)

	val1 := m["a"]
	fmt.Println("value of key a:", val1)

	fmt.Println("length of map:", len(m))

	delete(m, "b") // Deletes the key

	fmt.Println("After deletion,map:", m)

	_, prs := m["b"]
	fmt.Println("Presence of key b:", prs)

	mapWithInit := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Initialized Map :", mapWithInit)
}
