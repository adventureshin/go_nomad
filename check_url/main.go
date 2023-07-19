package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddis.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	var new_result requestResult
	for i:=0; i<len(urls); i++ {
		new_result= <-c
		results[new_result.url] = new_result.status
	}
	for url, status := range results{
		fmt.Println(url + status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >=400 {
		c <- requestResult{url:url, status: "Failed"}
	} else {
		c <- requestResult{url:url, status: "Ok"}
	}
}
