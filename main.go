package main

import (
  "fmt"
  "go_nomad/accounts"
  "log"
)


func main() {
  account := accounts.NewAccount("nico")
  fmt.Println(account)
  account.Deposit(10)
  fmt.Println(account)
  fmt.Println(account.Balance())
  err := account.Withdraw(5)
  if err != nil {
    fmt.Println(err)
    log.Fatalln(err)
  }
  account.ChangeOwner("lynn")
  fmt.Println(account)
  fmt.Println(account.Owner())
}
