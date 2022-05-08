package main

import (
  "fmt"
)

func main() {
  var h, w, r, c int
  fmt.Scanf("%d %d", &h, &w)
  fmt.Scanf("%d %d", &r, &c)

  if h == 1 && w == 1 {
    fmt.Printf("0\n")
  } else if h == 1 && w == 2 || h == 2 && w == 1 {
    fmt.Printf("1\n")
  } else if h == 1 && w >= 3 {
    if c == 1 || c == w {
      fmt.Printf("1\n")
    } else {
      fmt.Printf("2\n")
    }
  } else if h >= 3 && w == 1 {
    if r == 1 || r == h {
      fmt.Printf("1\n")
    } else {
      fmt.Printf("2\n")
    }
  } else {
    if r > 1 && r < h && c > 1 && c < w {
      fmt.Printf("4\n")
    } else if r == 1 && c == 1 || r == 1 && c == w || r == h && c == 1 || r == h && c == w {
      fmt.Printf("2\n")
    } else {
      fmt.Printf("3\n")
    }
  }
}
