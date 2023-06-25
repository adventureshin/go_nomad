package main

import (
  "fmt"
  "go_nomad/accounts"
)


func main() {
  account := accounts.NewAccount("nico")
  fmt.Println(account)
  account.Deposit(10)
  fmt.Println(account)
}
