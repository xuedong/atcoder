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

  counts := make(map[int]int)
  k := 0
  for i := 0; i < n; i++ {
    if _, ok := counts[nums[i]]; ok {
      counts[nums[i]] += 1
    } else {
      counts[nums[i]] = 1
      k++
    }
  }

  if k <= 2 {
    fmt.Println(0)
  } else {
    m := len(counts)
    values := make([]int, 0, m)

	for  _, value := range counts {
   	  values = append(values, value)
	}

    suffixes := make([]int, m-2)
    suffix := 0
    for i := m-1; i > 1; i-- {
      suffix += values[i]
      suffixes[i-2] = suffix
    }

    ans := 0
    for i1 := 0; i1 < m; i1++ {
      for i2 := i1+1; i2 < m-1; i2++ {
        ans += values[i1] * values[i2] * suffixes[i2-1]
      }
    }

    fmt.Printf("%d\n", ans)
  }
}
