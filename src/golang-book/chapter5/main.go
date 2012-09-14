package main

import "fmt"

// for and switch

func main() {
  for i := 0; i <= 10; i++ {
    
    if i % 2 == 0 {
      fmt.Printf("%d is even\n", i)
    
    } else if i == 7 { 
      fmt.Printf("%d is SEVEN\n", i)
    
    } else {
      fmt.Printf("%d is odd\n", i)
    }

    switch i {
      case 0: fmt.Println("oh, a zero!")
      case 10: fmt.Println("this is the end...")
      default: fmt.Println("nuttin' special")
    }
  }


}


