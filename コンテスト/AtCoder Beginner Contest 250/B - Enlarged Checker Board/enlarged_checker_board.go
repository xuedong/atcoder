package main

import (
  "fmt"
)

func main() {
  var n, a, b int
  fmt.Scanf("%d %d %d", &n, &a, &b)

  for i := 0; i < n*a; i++ {
    for j := 0; j < n*b; j++ {
      if (i/a + j/b) % 2 == 0 {
        fmt.Printf(".")
      } else {
        fmt.Printf("#")
      }
    }
    fmt.Println()
  }
}
