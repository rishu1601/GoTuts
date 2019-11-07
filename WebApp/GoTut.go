package main

import (
	"fmt"
	"net/http"
)

// There are no classes in go
//Instead we use * structs *

const usixteenmax float64 = 65535
const kmConverter float64 = 1.60934

type car struct {
	gasPedal      uint16
	brakePedal    uint16
	steeringWheel int16
	topSpeed      float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeed / usixteenmax)
}
func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeed / usixteenmax / kmConverter)
}

//Pointer receivers
func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeed = newSpeed
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa! Writing through responseWriter")
}
func main() {
	aCar := car{
		gasPedal:      52345,
		brakePedal:    0,
		steeringWheel: -10,
		topSpeed:      200.1,
	}
	// aCar.attributeName would be used to get the data of attribute
	fmt.Println(aCar)
	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	//Changing top speed through pointer receiver
	aCar.newTopSpeed(300.2)
	//After changing top speed
	fmt.Println("After changing top speed")
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
