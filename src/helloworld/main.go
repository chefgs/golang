package main

import (
  "fmt"
)

func main() {
  val := print()
  fmt.Println(val)
}

func print() string {
  str := "Hello World!"
  return str
}
