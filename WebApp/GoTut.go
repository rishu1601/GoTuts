package main

import (
	"fmt"
	"net/http"
)

// There are no classes in go
//Instead we use * structs *

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa! Writing through responseWriter")
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
