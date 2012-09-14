package main

import "fmt"

func main() {

  // an array
  var x [5] int
  
  for i :=  0; i < len(x); i++ {
    x[i] = 100 - i
  }

  fmt.Println(x)

  // a slice
  y := [] string {
    "muchos",
    "elementos",
    "por",
    "favor", // mandatory :)
  }

  fmt.Println(y)

  // a map
  z := map [string] int {
    "nil" : 0,
    "one" : 1,
    "two" : 2,
    "three" : 3,
  }
  delete(z, "nil")

 fmt.Println(z)
}

