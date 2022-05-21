package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scanf("%d", &n)

  strings := make([]string, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&strings[i])
  }

  counts := make([][]int, 10)
  for i := range counts {
    counts[i] = make([]int, 10)
  }

  for i := 0; i < n; i++ {
    for j := 0; j < 10; j++ {
	  counts[strings[i][j] - '0'][j]++
	}
  }

  ans := 10000
  for i := 0; i < 10; i++ {
    curr := 0
    for j := 0; j < 10; j++ {
      count := 10 * (counts[i][j] - 1) + j
      if count > curr {
        curr = count
      }
    }

    if curr < ans {
      ans = curr
    }
  }

  fmt.Printf("%d\n", ans)
}
