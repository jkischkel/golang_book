package main

import "fmt"

const RATIO float64 = 0.3048

func main() {
  fmt.Print("Feet: ")
  
  var input float64
  fmt.Scanf("%f", &input)

  output := input * RATIO
  fmt.Printf("%.2fft = %.2fm", input, output)
}

