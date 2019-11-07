package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//Multi Line print
	fmt.Fprintf(w, `<h1>Hi,where is this from?</h1>
	<p>nvwe fewf ewf ewfewf</p>
	<p>efef efw fw efefewfe</p>
	`)
	fmt.Fprintf(w, "<h1>Whoa!</h1> Writing through responseWriter")
	fmt.Fprintf(w, "<p>But not this way</p>")
	fmt.Fprintf(w, "<p>Its very static</p>")
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}
