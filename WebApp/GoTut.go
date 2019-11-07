package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Accessing the internet with Go

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body) //Response is received in bytes
	stringBody := string(bytes)           //Converting bytes to string
	fmt.Println(stringBody)
	resp.Body.Close() //Close the resources used
}
