package main

import (
  "fmt"
  "strings"
)


func lenAndUpper(name string) (length int, uppercase string) {
  length = len(name)
  uppercase = strings.ToUpper(name)
  return
}

func main() {
  totalLength, upperName := lenAndUpper("nico")
  fmt.Println(totalLength, upperName)
}
