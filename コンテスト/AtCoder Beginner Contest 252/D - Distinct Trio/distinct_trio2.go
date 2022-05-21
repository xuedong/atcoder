package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scanf("%d", &n)

  nums := make([]int, n)
  for i := 0; i < n; i++ {
    fmt.Scan(&nums[i])
  }

  counts := make([]int, 200001)
  for i := 0; i < n; i++ {
    counts[nums[i]] += 1
  }

  for i := 1; i < 200001; i++ {
    counts[i] += counts[i-1]
  }

  ans := 0
  for i := 0; i < n; i++ {
    ans += counts[nums[i]-1] * (n - counts[nums[i]])
  }

  fmt.Printf("%d\n", ans)
}
