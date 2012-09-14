package main

import "fmt"

func main() {
  xs := [] float64 { 40, 30, 23, 49, 26 }

  fmt.Println(average(xs))
}

func average(in [] float64) float64 {
  var total float64 = 0.0

  for _, v := range in {
    total += v  
  }

  return total / float64(len(in))
}

