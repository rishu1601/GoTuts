package main

import (
	"fmt"
	"net/http"
)

// There are no classes in go
//Instead we use * structs *

type car struct {
	gasPedal      uint32
	brakePedal    uint32
	steeringWheel int32
	topSpeed      float64
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa! Writing through responseWriter")
}
func main() {
	aCar := car{
		gasPedal:      123,
		brakePedal:    0,
		steeringWheel: -10,
		topSpeed:      200.1,
	}
	// aCar.attributeName would be used to get the data of attribute
	fmt.Println(aCar)
	fmt.Println(aCar.gasPedal)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
