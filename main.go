package main

import "fmt"

func main() {
  const name string = "nico"
  var name2 string = "lynn"
  name2 = "nico"
  fmt.Println(name, name2)
  // 축약형 but, 새로운 변수를 만들 때만 사용 가능, 함수 안에서만 사용 가능
  name3 := "lynn"
  fmt.Println(name3)
}
