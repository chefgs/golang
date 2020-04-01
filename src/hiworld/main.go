// Simple Hi World program
package main

import (
  "fmt"
)

func hifun() string {
  str := "Go hiworld program"
  return str
}

func main() {
  val := hifun()
  fmt.Println(val)
  fmt.Printf("%s\n",val)
}
