package main

import "fmt"
import "os"

func main() {

  if len(os.Args) == 1 {
    fmt.Println("hello old world")
    fmt.Println("more test")
    os.Exit(0)
  }

  fmt.Println("hello old error")
  os.Exit(1)

}


