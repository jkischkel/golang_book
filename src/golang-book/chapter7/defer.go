package main

import "os"
import "fmt"
import "log"

func main() {
  f, _ := os.Open("main.go")
  defer f.Close() // executed at method end
  defer fmt.Println("This is the end")

  data := make([]byte, 12)
  count, e := f.Read(data)

  if e != nil {
    log.Fatal(e)
  }
   
  fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

