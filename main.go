package main

import (
  "fmt"
)

type person struct {
  name string
  age int
  favFood []string
}

func main() {
  nico := person{
    name: "nico",
    age: 18,
    favFood: []string{"kimchi", "ramen"}, 
  }
  fmt.Println(nico)
}
