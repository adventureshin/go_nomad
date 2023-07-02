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
    "https://academy.nomadcoders.co/",
  }
  cErr := make(chan error)
  var result = make(map[string]string)
  for _, url := range urls {
    go hitUrl(url, cErr)
  }
  for _, url := range urls {
    err := <- cErr
    if err != nil {
      result[url] = "FAILED"
    } else {
      result[url] = "OK"
    }
  }
  for url, status := range result {
    fmt.Println(url + ": " + status)
  }
}

func hitUrl(url string, cErr chan error) {
  fmt.Println("Checking:", url)
  resp, err := http.Get(url)
  if err != nil {
    cErr <- errRequestFailed
  } else if resp.StatusCode >= 400 {
    cErr <- errRequestFailed
  }
  cErr <- nil
}
