package main

import (
  "fmt"
)

func main() {
  var n, k int
  fmt.Scanf("%d %d", &n, &k)

  tastiness := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&tastiness[i])
  }

  dislikes := make([]int, k)
  for i := 0; i < k; i++ {
    fmt.Scan(&dislikes[i])
  }

  best := 0
  for i := 0; i < n; i++ {
    if tastiness[i] > best {
      best = tastiness[i]
    }
  }

  flag := false
  for i := 0; i < n; i++ {
    if tastiness[i] == best {
      for j := 0; j < k; j++ {
        if i+1 == dislikes[j] {
          flag = true
          break
        }
      }
    }
  }

  if flag {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
