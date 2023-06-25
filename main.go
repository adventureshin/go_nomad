package main

import (
  "fmt"
  "go_nomad/dict"
  "log"
)

func searchDict(word string, d dict.Dictionary) string {
  find_word, nill := d.Search(word)
  if nill != nil {
    log.Fatal(nill)
  }
  return find_word
}


func main() {
  dictionary := dict.Dictionary{"first": "First word"}
  fmt.Println(dictionary.Search("first"))
  fmt.Println(dictionary.Search("second"))
  fmt.Println(searchDict("first", dictionary))
  err := dictionary.Add("second", "Second word")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(dictionary)
  fmt.Println(searchDict("second", dictionary))
  dictionary.Update("first", "First word update")
  fmt.Println(dictionary)
}


