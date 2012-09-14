package main

import "fmt"

// this is a comment

func main() {
  msg := "Hello, " // short initialization
  var name string = "Jan" // explicit

  fmt.Println(msg + name)

  const MAGIC int = 5
  fmt.Printf("%U", MAGIC)
}

