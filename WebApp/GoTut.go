package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var stringBody string

//Accessing the internet with Go

//Get resp from a particular url
func getResp(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	return resp, err
}

//Convert received response to string
func convertByteToString(resp *http.Response) string {
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	return stringBody
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, `<h1>Hi</h1>`)
	fmt.Println(w, "Hi")
}

func main() {
	url := "https://www.washingtonpost.com/news-sitemaps/index.xml"
	resp, _ := getResp(url)
	stringBody := convertByteToString(resp)
	fmt.Println(stringBody)
	// http.HandleFunc("/", indexHandler)
	// http.ListenAndServe(":8000", nil)
	resp.Body.Close() //Close the resources used

}
