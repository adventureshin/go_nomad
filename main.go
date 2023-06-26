package main

import (
  "fmt"
  "net/http"
  "errors"
)

var errRequestFailed = errors.New("Request failed")

func main() {
  urls := []string{
    "https://www.airbnb.com/",
    "https://www.google.com/",
    "https://www.amazon.com/",
    "https://www.reddit.com/",
    "https://www.google.com/",
    "https://soundcloud.com/",
    "https://www.facebook.com/",
    "https://www.instagram.com/",
    "https://academy.nomadcoders.c/",
  }
  var result = make(map[string]string)
  for _, url := range urls {
    err := hitURL(url)
    if err != nil {
      result[url] = "FAILED"
    } else {
      result[url] = "OK"
    }
  }
  var url string
  var status string
  for url, status = range result {
    fmt.Println(url, status)
  }
}

func hitURL(url string) error {
  fmt.Println("Checking:", url)
  resp, err := http.Get(url)
  if err != nil {
    return errRequestFailed
  }
  if resp.StatusCode >= 400 {
    fmt
    return errRequestFailed
  }
  return nil
}
