package main

import "fmt"
import "os"

func main() {

  if len(os.Args) == 1 {
    fmt.Println("hello old world")
    os.Exit(0)
  }

  fmt.Println("hello error")
  os.Exit(1)

}


