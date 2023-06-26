package main

import (
  "fmt"
  "time"
)

type people struct {
  person string
  age int
}


func main() {
  c := make(chan bool)
  peopleList := [2]people {}
  peopleList[0] = people {"nico", 5}
  peopleList[1] = people {"flynn", 1}
  for _, person := range peopleList {
    fmt.Println("start: " + person.person)
    go isSexy(person.person, person.age, c)
  }
  fmt.Println(<-c)
  fmt.Println(<-c)
}

func isSexy(person string, delay int, c chan bool){
  time.Sleep(time.Second * time.Duration(delay))
  fmt.Println("send to channel:" + person)
  c <- true
}


