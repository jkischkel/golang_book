package main

import "fmt"

func main() {
  var number uint = 10
  result := fib(number)

  fmt.Printf("fib(%d) = %d\n", number, result)
}

func fib(n uint) uint {
  switch n {
    case 0:  return 0
    case 1:  return 1
  }

  return fib(n - 1) + fib(n - 2) 
}

